package non_terminal

import (
	"github.com/walkerrandolphsmith/parser/ast/boolean_expression"
)

type NonTerminal struct {
	Left boolean_expression.BooleanExpression
	Right boolean_expression.BooleanExpression
}

func (node *NonTerminal) SetLeft(left boolean_expression.BooleanExpression) {
	node.Left = left
}

func (node *NonTerminal) SetRight(right boolean_expression.BooleanExpression) {
	node.Right = right
}