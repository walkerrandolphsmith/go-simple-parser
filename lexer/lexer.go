package lexer

import (
	"regexp"
)

type Token int

const (
	INVALID  Token = -1
	NONE     Token = 0
	OR       Token = 1
	AND      Token = 2
	NOT      Token = 3
	TRUE     Token = 4
	FALSE    Token = 5
	LEFT     Token = 6
	RIGHT    Token = 7
)

type Lexer struct {
	tokens []string
	index int
}

func New(input string) Lexer {
	r := regexp.MustCompile(`\(|\)|true|false|and|or|not`)
	tokens := r.FindAllString(input, -1)
	return Lexer{
		tokens: tokens,
		index: -1,
	}
}

func (l *Lexer) NextSymbol() Token {
	l.index++

	value := ""
	if l.index < len(l.tokens) {
		value = l.tokens[l.index]
	}

	switch value {
	case "(":
		return LEFT
	case ")":
		return RIGHT
	case "and":
		return AND
	case "or":
		return OR
	case "not":
		return NOT
	case "true":
		return TRUE
	case "false":
		return FALSE
	default:
		return INVALID
	}
}