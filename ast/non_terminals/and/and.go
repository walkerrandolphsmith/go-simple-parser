package and

import (
	"github.com/walkerrandolphsmith/parser/ast/non_terminal"
)

type And struct {
	non_terminal.NonTerminal
}

func New() And {
	return And{non_terminal.NonTerminal{}}
}

func (node *And) Interpret() bool {
	return node.Left.Interpret() && node.Right.Interpret()
}

func (node *And) Stringify() string {
	return "(" + node.Left.Stringify() + " & "  + node.Right.Stringify() + ")"
}