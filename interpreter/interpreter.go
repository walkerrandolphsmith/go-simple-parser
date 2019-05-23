package interpreter

import (
	"github.com/walkerrandolphsmith/parser/lexer"
	"github.com/walkerrandolphsmith/parser/parser"
)

func Eval(expression string) bool {
	l := lexer.New(expression)
	p := parser.New(l)
	ast := p.Build()
	return ast.Interpret()
}
