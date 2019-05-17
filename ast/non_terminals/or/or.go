package or

import (
	"github.com/walkerrandolphsmith/parser/ast/non_terminal"
)

type Or struct {
	non_terminal.NonTerminal
}

func New() Or {
	return Or {non_terminal.NonTerminal{}}
}

func (node *Or) Interpret() bool {
	return node.Left.Interpret() || node.Right.Interpret()
}

func (node *Or) Stringify() string {
	return "(" + node.Left.Stringify() + " || " + node.Right.Stringify() + ")"
}