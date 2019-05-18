package terminal

import (
	"strconv"
	"github.com/walkerrandolphsmith/parser/ast/boolean_expression"
)

type Terminal struct {
	Value bool
	boolean_expression.BooleanExpression
}

func New(value bool) Terminal {
	return Terminal { Value: value }
}

func (t Terminal) Interpret() bool {
	return t.Value
}

func (t Terminal) Stringify() string {
	return strconv.FormatBool(t.Value)
}