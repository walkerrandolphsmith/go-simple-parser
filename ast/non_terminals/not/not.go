package not


import (
	"github.com/walkerrandolphsmith/parser/ast/non_terminal"
	"github.com/walkerrandolphsmith/parser/ast/boolean_expression"
)

type Not struct {
	non_terminal.NonTerminal
}

func New() Not {
	return Not{non_terminal.NonTerminal{}}
}

func (non_terminal *Not) SetRight(right boolean_expression.BooleanExpression) {
	// possibly throw?
}

func (node *Not) Interpret() bool {
	return !node.Left.Interpret()
}

func (node *Not) Stringify() string {
	return "!" + node.Left.Stringify()
}