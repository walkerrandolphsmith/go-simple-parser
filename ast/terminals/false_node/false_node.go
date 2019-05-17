package false_node

import (
	"github.com/walkerrandolphsmith/parser/ast/terminal"
)

type False struct {
	terminal.Terminal
}

func New() False {
	return False{terminal.Terminal{Value: false}}
}	