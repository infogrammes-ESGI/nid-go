package main

import (
	"bufio"
	"io"
	"unicode"
)

// TOKENS

type Token int

const (
	EOF = iota
	ACTION
	PROTOCOL
	LPAREN
	RPAREN
	COMA
	COLON
	SEMICOLON
	QUOTATION_MARK
)

var tokens = []string{
	EOF:            "EOF",
	ACTION:         "ACTION",
	PROTOCOL:       "PROTOCOL",
	LPAREN:         "(",
	RPAREN:         ")",
	COMA:           ",",
	COLON:          ":",
	SEMICOLON:      ";",
	QUOTATION_MARK: "\"",
}

// list of actions possible
var ACTIONS = []string{
	"alert",
	"log",
	"pass",
}

func (t Token) String() string {
	return tokens[t]
}

// LEXER

type Position struct {
	line   int
	column int
}

func (p *Position) nextLine() {
	p.line++
	p.column = 0
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Advance() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			// TODO: change this to a syslog alert
			panic(err)
		}

		switch r {
		case '\n':
			l.pos.nextLine()
			continue
		case '\t':
			continue
		case '(':
			return l.pos, LPAREN, "("
		case ')':
			return l.pos, RPAREN, ")"
		case ',':
			return l.pos, COMA, ","
		case ':':
			return l.pos, COLON, ":"
		case ';':
			return l.pos, SEMICOLON, ";"
		case '"':
			return l.pos, QUOTATION_MARK, "\""
		default:
			if unicode.IsSpace(r) {
				continue
			}
		}

		l.pos.column++
	}
}
