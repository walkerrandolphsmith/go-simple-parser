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

	node.SetRight(trueNode)

	if node.Right != nil {
		t.Errorf("Uranary opertators can't have a Right node")
	}

	if node.Stringify() != "!false" {
		t.Errorf("Not not stringified correctly")
	}
}