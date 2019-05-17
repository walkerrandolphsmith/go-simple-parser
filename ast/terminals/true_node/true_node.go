package true_node

import (
	"github.com/walkerrandolphsmith/parser/ast/terminal"
)

type True struct {
	terminal.Terminal
}

func New() True {
	return True{terminal.Terminal{Value: true}}
}