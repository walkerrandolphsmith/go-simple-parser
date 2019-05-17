package not

import (
	"testing"
	"github.com/walkerrandolphsmith/parser/ast/terminals/true_node"
	"github.com/walkerrandolphsmith/parser/ast/terminals/false_node"
)

func TestInterpret(t *testing.T) {
	node := New()
	trueNode := true_node.New()
	falseNode := false_node.New()

	node.SetLeft(trueNode)
	
	if node.Interpret() != false {
		t.Errorf("not true != false")
	}

	node.SetLeft(falseNode)
	
	if node.Interpret() != true {
		t.Errorf("not false != true")
	}
}