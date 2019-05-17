package parser

import (
	"github.com/walkerrandolphsmith/parser/ast/boolean_expression"
	"github.com/walkerrandolphsmith/parser/ast/terminals/true_node"
	"github.com/walkerrandolphsmith/parser/ast/terminals/false_node"
)

type RecursiveDescentParser struct {
	symbol int
	root boolean_expression.BooleanExpression
	t true_node.True
	f false_node.False
}