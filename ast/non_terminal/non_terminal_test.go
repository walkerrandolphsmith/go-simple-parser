package non_terminal

import (
	"testing"
	"github.com/walkerrandolphsmith/parser/ast/terminal"
)
func TestSetLeft(t *testing.T) {
	node := NonTerminal{}
	child := terminal.New(true)
	node.SetLeft(child)
	if node.Left == nil {
		t.Errorf("Left node is not set");
	}
	if node.Left.Interpret() != true {
		t.Errorf("Left was not set to the correct node")
	}
}

func TestSetRight(t *testing.T) {
	node := NonTerminal{}
	child := terminal.New(false)
	node.SetRight(child)
	if node.Right == nil {
		t.Errorf("Right node is not set");
	}
	if node.Right.Interpret() != false {
		t.Errorf("Right was not set to the correct node")
	}
}