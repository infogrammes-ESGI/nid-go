// Code generated by goyacc -p RuleParser -o rule-parser.go rule-parser.y. DO NOT EDIT.

//line rule-parser.y:3

package main

import __yyfmt__ "fmt"

//line rule-parser.y:4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

// it validates simple ip address (e.g. '192.168.1.1') so we need to be careful
var IPV4_RANGE_REGEX = regexp.MustCompile(`^((([1-9]?\d|[12]\d\d)\.){3}([1-9]?\d|[12]\d\d))?-(((([1-9]?\d|[12]\d\d)\.){3}([1-9]?\d|[12]\d\d)))?$`)

var IPV4_REGEX = regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`)
var IPV6_REGEX = regexp.MustCompile(`((([0-9a-fA-F]){1,4})\\:){7}([0-9a-fA-F]){1,4}`)
var PORT_REGEX = regexp.MustCompile(`(?:[1-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])`)
var PORT_RANGE_REGEX = regexp.MustCompile(`^(?:[1-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])?-(?:[1-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])?$`)
var INTEGER_REGEX = regexp.MustCompile(`[0-9]+`)

//line rule-parser.y:38
type RuleParserSymType struct {
	yys         int
	integer_val int
	string_val  string
}

const ACTION = 57346
const ACTION_IDENTIFIER = 57347
const PROTO_IDENTIFIER = 57348
const IDENTIFIER = 57349
const INTEGER = 57350
const SLASH = 57351
const DOLLAR = 57352
const ASTERISK = 57353
const MINUS = 57354
const LPAREN = 57355
const RPAREN = 57356
const ARROW = 57357
const COLON = 57358
const QUOTE = 57359
const COMMA = 57360
const SEMICOLON = 57361
const LCOMMENT = 57362
const RCOMMENT = 57363
const ANY_KEYWORD = 57364
const IPV4_ADDR = 57365
const IPV4_RANGE = 57366
const IPV6_ADDR = 57367
const IPV6_RANGE = 57368
const PORT_RANGE = 57369
const UMINUS = 57370

var RuleParserToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ACTION",
	"ACTION_IDENTIFIER",
	"PROTO_IDENTIFIER",
	"IDENTIFIER",
	"INTEGER",
	"SLASH",
	"DOLLAR",
	"ASTERISK",
	"MINUS",
	"LPAREN",
	"RPAREN",
	"ARROW",
	"COLON",
	"QUOTE",
	"COMMA",
	"SEMICOLON",
	"LCOMMENT",
	"RCOMMENT",
	"ANY_KEYWORD",
	"IPV4_ADDR",
	"IPV4_RANGE",
	"IPV6_ADDR",
	"IPV6_RANGE",
	"PORT_RANGE",
	"'|'",
	"'&'",
	"'+'",
	"'%'",
	"UMINUS",
	"'('",
	"')'",
	"'-'",
}

var RuleParserStatenames = [...]string{}

const RuleParserEofCode = 1
const RuleParserErrCode = 2
const RuleParserInitialStackSize = 16

//line rule-parser.y:149

/* START OF GOLANG CODE */

/*
From: https://pkg.go.dev/golang.org/x/tools/cmd/goyacc

	type yyLexer interface {
		Lex(lval *yySymType) int
		Error(e string)
	}
*/
type RuleParserLex struct {
	s   string
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

//line yacctab:1
var RuleParserExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const RuleParserPrivate = 57344

const RuleParserLast = 21

var RuleParserAct = [...]int8{
	11, 12, 8, 9, 6, 10, 7, 5, 4, 3,
	2, 1, 0, 0, 0, 0, 0, 14, 0, 0,
	13,
}

var RuleParserPact = [...]int16{
	-1000, 4, -1000, 2, -20, -22, -1000, -1000, -1000, -14,
	-1000, -1000, -20, -22, -1000,
}

var RuleParserPgo = [...]int8{
	0, 12, 12, 12, 3, 7, 11, 10,
}

var RuleParserR1 = [...]int8{
	0, 6, 6, 7, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 5, 5, 5, 4,
	4, 3,
}

var RuleParserR2 = [...]int8{
	0, 2, 0, 7, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 1, 2, 1, 1, 1, 1,
	1, 1,
}

var RuleParserChk = [...]int16{
	-1000, -6, -7, 5, 6, -5, 24, 26, 22, -4,
	27, 22, 15, -5, -4,
}

var RuleParserDef = [...]int8{
	2, -2, 1, 0, 0, 0, 16, 17, 18, 0,
	19, 20, 0, 0, 3,
}

var RuleParserTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 31, 29, 3,
	33, 34, 3, 30, 3, 35, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 28,
}

var RuleParserTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 32,
}

var RuleParserTok3 = [...]int8{
	0,
}

var RuleParserErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	RuleParserDebug        = 0
	RuleParserErrorVerbose = false
)

type RuleParserLexer interface {
	Lex(lval *RuleParserSymType) int
	Error(s string)
}

type RuleParserParser interface {
	Parse(RuleParserLexer) int
	Lookahead() int
}

type RuleParserParserImpl struct {
	lval  RuleParserSymType
	stack [RuleParserInitialStackSize]RuleParserSymType
	char  int
}

func (p *RuleParserParserImpl) Lookahead() int {
	return p.char
}

func RuleParserNewParser() RuleParserParser {
	return &RuleParserParserImpl{}
}

const RuleParserFlag = -1000

func RuleParserTokname(c int) string {
	if c >= 1 && c-1 < len(RuleParserToknames) {
		if RuleParserToknames[c-1] != "" {
			return RuleParserToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func RuleParserStatname(s int) string {
	if s >= 0 && s < len(RuleParserStatenames) {
		if RuleParserStatenames[s] != "" {
			return RuleParserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func RuleParserErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !RuleParserErrorVerbose {
		return "syntax error"
	}

	for _, e := range RuleParserErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + RuleParserTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(RuleParserPact[state])
	for tok := TOKSTART; tok-1 < len(RuleParserToknames); tok++ {
		if n := base + tok; n >= 0 && n < RuleParserLast && int(RuleParserChk[int(RuleParserAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if RuleParserDef[state] == -2 {
		i := 0
		for RuleParserExca[i] != -1 || int(RuleParserExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; RuleParserExca[i] >= 0; i += 2 {
			tok := int(RuleParserExca[i])
			if tok < TOKSTART || RuleParserExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if RuleParserExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += RuleParserTokname(tok)
	}
	return res
}

func RuleParserlex1(lex RuleParserLexer, lval *RuleParserSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(RuleParserTok1[0])
		goto out
	}
	if char < len(RuleParserTok1) {
		token = int(RuleParserTok1[char])
		goto out
	}
	if char >= RuleParserPrivate {
		if char < RuleParserPrivate+len(RuleParserTok2) {
			token = int(RuleParserTok2[char-RuleParserPrivate])
			goto out
		}
	}
	for i := 0; i < len(RuleParserTok3); i += 2 {
		token = int(RuleParserTok3[i+0])
		if token == char {
			token = int(RuleParserTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(RuleParserTok2[1]) /* unknown char */
	}
	if RuleParserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", RuleParserTokname(token), uint(char))
	}
	return char, token
}

func RuleParserParse(RuleParserlex RuleParserLexer) int {
	return RuleParserNewParser().Parse(RuleParserlex)
}

func (RuleParserrcvr *RuleParserParserImpl) Parse(RuleParserlex RuleParserLexer) int {
	var RuleParsern int
	var RuleParserVAL RuleParserSymType
	var RuleParserDollar []RuleParserSymType
	_ = RuleParserDollar // silence set and not used
	RuleParserS := RuleParserrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	RuleParserstate := 0
	RuleParserrcvr.char = -1
	RuleParsertoken := -1 // RuleParserrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		RuleParserstate = -1
		RuleParserrcvr.char = -1
		RuleParsertoken = -1
	}()
	RuleParserp := -1
	goto RuleParserstack

ret0:
	return 0

ret1:
	return 1

RuleParserstack:
	/* put a state and value onto the stack */
	if RuleParserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", RuleParserTokname(RuleParsertoken), RuleParserStatname(RuleParserstate))
	}

	RuleParserp++
	if RuleParserp >= len(RuleParserS) {
		nyys := make([]RuleParserSymType, len(RuleParserS)*2)
		copy(nyys, RuleParserS)
		RuleParserS = nyys
	}
	RuleParserS[RuleParserp] = RuleParserVAL
	RuleParserS[RuleParserp].yys = RuleParserstate

RuleParsernewstate:
	RuleParsern = int(RuleParserPact[RuleParserstate])
	if RuleParsern <= RuleParserFlag {
		goto RuleParserdefault /* simple state */
	}
	if RuleParserrcvr.char < 0 {
		RuleParserrcvr.char, RuleParsertoken = RuleParserlex1(RuleParserlex, &RuleParserrcvr.lval)
	}
	RuleParsern += RuleParsertoken
	if RuleParsern < 0 || RuleParsern >= RuleParserLast {
		goto RuleParserdefault
	}
	RuleParsern = int(RuleParserAct[RuleParsern])
	if int(RuleParserChk[RuleParsern]) == RuleParsertoken { /* valid shift */
		RuleParserrcvr.char = -1
		RuleParsertoken = -1
		RuleParserVAL = RuleParserrcvr.lval
		RuleParserstate = RuleParsern
		if Errflag > 0 {
			Errflag--
		}
		goto RuleParserstack
	}

RuleParserdefault:
	/* default state action */
	RuleParsern = int(RuleParserDef[RuleParserstate])
	if RuleParsern == -2 {
		if RuleParserrcvr.char < 0 {
			RuleParserrcvr.char, RuleParsertoken = RuleParserlex1(RuleParserlex, &RuleParserrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if RuleParserExca[xi+0] == -1 && int(RuleParserExca[xi+1]) == RuleParserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			RuleParsern = int(RuleParserExca[xi+0])
			if RuleParsern < 0 || RuleParsern == RuleParsertoken {
				break
			}
		}
		RuleParsern = int(RuleParserExca[xi+1])
		if RuleParsern < 0 {
			goto ret0
		}
	}
	if RuleParsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			RuleParserlex.Error(RuleParserErrorMessage(RuleParserstate, RuleParsertoken))
			Nerrs++
			if RuleParserDebug >= 1 {
				__yyfmt__.Printf("%s", RuleParserStatname(RuleParserstate))
				__yyfmt__.Printf(" saw %s\n", RuleParserTokname(RuleParsertoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for RuleParserp >= 0 {
				RuleParsern = int(RuleParserPact[RuleParserS[RuleParserp].yys]) + RuleParserErrCode
				if RuleParsern >= 0 && RuleParsern < RuleParserLast {
					RuleParserstate = int(RuleParserAct[RuleParsern]) /* simulate a shift of "error" */
					if int(RuleParserChk[RuleParserstate]) == RuleParserErrCode {
						goto RuleParserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if RuleParserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", RuleParserS[RuleParserp].yys)
				}
				RuleParserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if RuleParserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", RuleParserTokname(RuleParsertoken))
			}
			if RuleParsertoken == RuleParserEofCode {
				goto ret1
			}
			RuleParserrcvr.char = -1
			RuleParsertoken = -1
			goto RuleParsernewstate /* try again in the same state */
		}
	}

	/* reduction by production RuleParsern */
	if RuleParserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", RuleParsern, RuleParserStatname(RuleParserstate))
	}

	RuleParsernt := RuleParsern
	RuleParserpt := RuleParserp
	_ = RuleParserpt // guard against "declared and not used"

	RuleParserp -= int(RuleParserR2[RuleParsern])
	// RuleParserp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if RuleParserp+1 >= len(RuleParserS) {
		nyys := make([]RuleParserSymType, len(RuleParserS)*2)
		copy(nyys, RuleParserS)
		RuleParserS = nyys
	}
	RuleParserVAL = RuleParserS[RuleParserp+1]

	/* consult goto table to find next state */
	RuleParsern = int(RuleParserR1[RuleParsern])
	RuleParserg := int(RuleParserPgo[RuleParsern])
	RuleParserj := RuleParserg + RuleParserS[RuleParserp].yys + 1

	if RuleParserj >= RuleParserLast {
		RuleParserstate = int(RuleParserAct[RuleParserg])
	} else {
		RuleParserstate = int(RuleParserAct[RuleParserj])
		if int(RuleParserChk[RuleParserstate]) != -RuleParsern {
			RuleParserstate = int(RuleParserAct[RuleParserg])
		}
	}
	// dummy call; replaced with literal code
	switch RuleParsernt {

	case 3:
		RuleParserDollar = RuleParserS[RuleParserpt-7 : RuleParserpt+1]
//line rule-parser.y:90
		{
			var new_rule = Rule{}
			new_rule.action = RuleParserDollar[1].string_val
			new_rule.protocol = RuleParserDollar[2].string_val
			new_rule.in_network = RuleParserDollar[3].string_val
			new_rule.in_ports = RuleParserDollar[4].string_val
			new_rule.out_network = RuleParserDollar[6].string_val
			new_rule.out_ports = RuleParserDollar[7].string_val
			Add_New_Rule(new_rule)
			fmt.Printf("RULES = %v\n", Get_rules())
		}
	case 4:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:104
		{
			RuleParserVAL.integer_val = RuleParserDollar[2].integer_val
		}
	case 5:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:105
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val + RuleParserDollar[3].integer_val
		}
	case 6:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:106
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val - RuleParserDollar[3].integer_val
		}
	case 7:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:107
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val * RuleParserDollar[3].integer_val
		}
	case 8:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:108
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val / RuleParserDollar[3].integer_val
		}
	case 9:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:109
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val % RuleParserDollar[3].integer_val
		}
	case 10:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:110
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val & RuleParserDollar[3].integer_val
		}
	case 11:
		RuleParserDollar = RuleParserS[RuleParserpt-3 : RuleParserpt+1]
//line rule-parser.y:111
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val | RuleParserDollar[3].integer_val
		}
	case 12:
		RuleParserDollar = RuleParserS[RuleParserpt-2 : RuleParserpt+1]
//line rule-parser.y:112
		{
			RuleParserVAL.integer_val = -RuleParserDollar[2].integer_val
		}
	case 14:
		RuleParserDollar = RuleParserS[RuleParserpt-1 : RuleParserpt+1]
//line rule-parser.y:118
		{
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val
			if RuleParserDollar[1].integer_val == 0 {
				base = 8
			} else {
				base = 10
			}
		}
	case 15:
		RuleParserDollar = RuleParserS[RuleParserpt-2 : RuleParserpt+1]
//line rule-parser.y:127
		{
			RuleParserVAL.integer_val = base*RuleParserDollar[1].integer_val + RuleParserDollar[2].integer_val
		}
	case 21:
		RuleParserDollar = RuleParserS[RuleParserpt-1 : RuleParserpt+1]
//line rule-parser.y:141
		{
			if RuleParserDollar[1].integer_val > 65535 || RuleParserDollar[1].integer_val < 1 {
				// TODO: throw error
			}
			RuleParserVAL.integer_val = RuleParserDollar[1].integer_val
		}
	}
	goto RuleParserstack /* stack new state and value */
}
