package parser

import (
	"github.com/walkerrandolphsmith/parser/ast/boolean_expression"
	"github.com/walkerrandolphsmith/parser/ast/terminals/true_node"
	"github.com/walkerrandolphsmith/parser/ast/terminals/false_node"
	"github.com/walkerrandolphsmith/parser/ast/non_terminals/not"
	"github.com/walkerrandolphsmith/parser/ast/non_terminals/and"
	"github.com/walkerrandolphsmith/parser/ast/non_terminals/or"
	"github.com/walkerrandolphsmith/parser/lexer"
)

type Parser struct {
	symbol lexer.Token
	root boolean_expression.BooleanExpression
	l lexer.Lexer
}

func New(lex lexer.Lexer) Parser {
	return Parser{l: lex}
}

func (parser *Parser) expression() {
	parser.term()
	for parser.symbol == lexer.OR {
		node := or.New()
		node.SetLeft(parser.root)
		parser.term()
		node.SetRight(parser.root)
		parser.root = &node
	}
}

func (parser *Parser) term() {
	parser.factor()
	for parser.symbol == lexer.AND {
		node := and.New()
		node.SetLeft(parser.root)
		parser.factor()
		node.SetRight(parser.root)
		parser.root = &node
	}
}

func (parser *Parser) factor() {
	symbol := parser.l.NextSymbol()
	if symbol == lexer.TRUE {
		parser.root = true_node.New()
		parser.symbol = parser.l.NextSymbol()
	} else if symbol == lexer.FALSE {
		parser.root = false_node.New()
		parser.symbol = parser.l.NextSymbol();
	} else if symbol == lexer.NOT {
		node := not.New()
		parser.factor();
		node.SetLeft(parser.root)
		parser.root = &node
	} else if symbol == lexer.LEFT {
		parser.expression()
		parser.symbol = parser.l.NextSymbol()
	} else {
		//throw new
	}
}

func (parser *Parser) Build() boolean_expression.BooleanExpression {
	parser.expression()
	return parser.root
}