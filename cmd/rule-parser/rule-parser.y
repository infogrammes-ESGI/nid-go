// golang headers
%{

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func array_contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

var base int

// it validates simple ip address (e.g. '192.168.1.1') so we need to be careful
var IPV4_RANGE_REGEX = regexp.MustCompile(`^((([1-9]?\d|[12]\d\d)\.){3}([1-9]?\d|[12]\d\d))?-(((([1-9]?\d|[12]\d\d)\.){3}([1-9]?\d|[12]\d\d)))?$`)

var IPV4_REGEX = regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`)
var IPV6_REGEX = regexp.MustCompile(`((([0-9a-fA-F]){1,4})\\:){7}([0-9a-fA-F]){1,4}`)
var PORT_REGEX = regexp.MustCompile(`(?:[1-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])`)
var PORT_RANGE_REGEX = regexp.MustCompile(`^(?:[1-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])?-(?:[1-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])?$`)
var INTEGER_REGEX = regexp.MustCompile(`[0-9]+`)
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	integer_val int
	string_val string
}

// tokens
%token ACTION
%token <string_val> ACTION_IDENTIFIER
%token <string_val> PROTO_IDENTIFIER
%token <string_val> IDENTIFIER
%token <integer_val> INTEGER
%token SLASH
%token DOLLAR
%token ASTERISK
%token MINUS
%token LPAREN
%token RPAREN
%token ARROW
%token COLON
%token QUOTE
%token COMMA
%token SEMICOLON
%token LCOMMENT
%token RCOMMENT
%token <string_val> ANY_KEYWORD
%token <string_val> IPV4_ADDR
%token <string_val> IPV4_RANGE
%token <string_val> IPV6_ADDR
%token <string_val> IPV6_RANGE
%token <string_val> PORT_RANGE


// functions' type
%type <integer_val> expr number port_number
%type <string_val> port_range network_range

// arithmetic logic
%left '|'
%left '&'
%left '+'  MINUS
%left ASTERISK  SLASH  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */


/* START OF GRAMMAR */
%%

rules:		rules rule
    	|	/* nothing */
	;


rule:		ACTION_IDENTIFIER PROTO_IDENTIFIER network_range port_range ARROW network_range port_range {
			var new_rule = Rule{}
			new_rule.action = $1
			new_rule.protocol = $2
			new_rule.in_network = $3
			new_rule.in_ports = $4
			new_rule.out_network = $6
			new_rule.out_ports = $7
			Add_New_Rule(new_rule)
			fmt.Printf("RULES = %v\n", Get_rules())
		}
	;


expr:		'(' expr ')'		{ $$  =  $2 }
	|	expr '+' expr		{ $$  =  $1 + $3 }
	|	expr MINUS expr		{ $$  =  $1 - $3 }
	|	expr ASTERISK expr		{ $$  =  $1 * $3 }
	|	expr SLASH expr		{ $$  =  $1 / $3 }
	|	expr '%' expr		{ $$  =  $1 % $3 }
	|	expr '&' expr		{ $$  =  $1 & $3 }
	|	expr '|' expr		{ $$  =  $1 | $3 }
	|	'-'  expr %prec UMINUS	{ $$  = -$2  }
	//|	IDENTIFIER			{ $$  = regs[$1] }
	|	number
	;

number:		INTEGER
		{
			$$ = $1;
			if $1==0 {
				base = 8
			} else {
				base = 10
			}
		}

	|    number INTEGER		{ $$ = base * $1 + $2 }
	;


network_range: 	IPV4_RANGE
		|	IPV6_RANGE
		|	ANY_KEYWORD
		;

port_range: 	PORT_RANGE
		|	ANY_KEYWORD
		;


port_number: number {
		if $1 > 65535 || $1 < 1 {
			// TODO: throw error
		}
		$$ = $1
	}
	;

%%

/* START OF GOLANG CODE */


/*
From: https://pkg.go.dev/golang.org/x/tools/cmd/goyacc
type yyLexer interface {
	Lex(lval *yySymType) int
	Error(e string)
}
*/
type RuleParserLex struct {
	s string
	pos int
}

func (l *RuleParserLex) read_until(characters string) string {
	/*
	Read until one of the char in the 'characters' parameter is reached.
	*/
	var res = make([]rune, 0)

	for !strings.ContainsRune(characters, rune(l.s[l.pos])) {
		res = append(res, rune(l.s[l.pos]))
		l.pos += 1

		if l.pos == len(l.s) {
			return string(res)
		}
	}
	return string(res)
}

func (l *RuleParserLex) Lex(lval *RuleParserSymType) int {
	var c rune
	for {
		if l.pos == len(l.s) {
			return 0
		}

		c = rune(l.s[l.pos])
		if c == ' ' || c == '\n' {
			l.pos += 1
			continue
		}
		break
	}
	
	var token string = l.read_until("\n \t")
	
	if token == "(" {
		return LPAREN
	} else if token == ")" {
		return RPAREN
	} else if token == "\"" {
		return QUOTE
	} else if token == "," {
		return COMMA
	} else if token == "$" {
		return DOLLAR
	} else if token == ";" {
		return SEMICOLON
	} else if token == "/" {
		return SLASH
	} else if token == "-" {
		return MINUS
	} else if token == "->" {
		return ARROW
	} else if token == "*" {
		return ASTERISK
	} else if token == "any" {
		lval.string_val = "any" // need to specify it because we will store it as port or network range in the final struct
		return ANY_KEYWORD
	} else if array_contains(LIST_PROTOCOLS, token) {
		lval.string_val = token
		return PROTO_IDENTIFIER
	} else if array_contains(LIST_ACTIONS, token) {
		lval.string_val = token
		return ACTION_IDENTIFIER
	} else if PORT_RANGE_REGEX.Match([]byte(token)) {
		lval.string_val = token
		return PORT_RANGE
	} else if IPV4_REGEX.Match([]byte(token)) {
		lval.string_val = token
		return IPV4_ADDR
	} else if IPV4_RANGE_REGEX.Match([]byte(token)) {
		lval.string_val = token
		return IPV4_RANGE
	} else if IPV6_REGEX.Match([]byte(token)) {
		lval.string_val = token
		return IPV6_ADDR
	} else if INTEGER_REGEX.Match([]byte(token)) {
		lval.integer_val, _ = strconv.Atoi(token)
		return INTEGER
	} else {
		lval.string_val = token
		return IDENTIFIER
	}

	return int(c)
}

func (l *RuleParserLex) Error(s string) {
	// TODO: syslog the error
	fmt.Printf("Rule-Parser error: %s\n", s)
}


func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("Rule: ")
		if eqn, ok = readline(fi); ok {
			RuleParserParse(&RuleParserLex{s: eqn})
		} else {
			break
		}
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
