// golang headers
%{

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strings"
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
%token LPAREN
%token RPAREN
%token ARROW
%token COLON
%token QUOTE
%token COMMA
%token SEMICOLON
%token LCOMMENT
%token RCOMMENT
%token ANY_KEYWORD
%token <string_val> IP_ADDR


// functions' type
%type <integer_val> expr number port_number
%type <string_val> port_range network_range

// arithmetic logic
%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */


/* START OF GRAMMAR */
%%

rules:		rules rule '\n'
    	|	/* nothing */
	;


rule:		ACTION_IDENTIFIER PROTO_IDENTIFIER {
			var new_rule = Rule{}
			new_rule.action = $1
			new_rule.protocol = $2
			Add_New_Rule(new_rule)
			fmt.Printf("RULES = %v\n", Get_rules())
		}
	;


expr:		'(' expr ')'		{ $$  =  $2 }
	|	expr '+' expr		{ $$  =  $1 + $3 }
	|	expr '-' expr		{ $$  =  $1 - $3 }
	|	expr '*' expr		{ $$  =  $1 * $3 }
	|	expr '/' expr		{ $$  =  $1 / $3 }
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


network_range: IP_ADDR {
				$$ = $1
			}
			;

port_range: port_number {
				$$ = string($1)
			}
		|	port_number '-' {
				$$ = string($1) + "-"
		}
		|	'-' port_number {
				$$ = "-" + string($2)
		}
		|	port_number '-' port_number{
				$$ = string($1) + "-" + string($3)
		}
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
		if c == ' ' {
			l.pos += 1
			continue
		}
		break
	}

	if unicode.IsDigit(c) {
		lval.integer_val = int(c) - '0'
		return INTEGER
	} else if unicode.IsLetter(c) {
		lval.string_val = l.read_until("\n \t")

		if array_contains(LIST_PROTOCOLS, lval.string_val) {
			return PROTO_IDENTIFIER
		} else if array_contains(LIST_ACTIONS, lval.string_val) {
			return ACTION_IDENTIFIER
		} else if lval.string_val == "any" {
			return ANY_KEYWORD
		}
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
