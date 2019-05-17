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
	t := Terminal { Value: value }
	return t
}

func (t Terminal) Interpret() bool {
	return t.Value
}

func Stringify(t Terminal) string {
	return strconv.FormatBool(t.Value)
}