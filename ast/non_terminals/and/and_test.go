package and

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
	node.SetRight(trueNode)
	
	if node.Interpret() != true {
		t.Errorf("true and true != true")
	}

	node.SetLeft(trueNode)
	node.SetRight(falseNode)
	
	if node.Interpret() != false {
		t.Errorf("true and false != false")
	}

	node.SetLeft(falseNode)
	node.SetRight(trueNode)
	
	if node.Interpret() != false {
		t.Errorf("false and true != false")
	}

	node.SetLeft(falseNode)
	node.SetRight(falseNode)
	
	if node.Interpret() != false {
		t.Errorf("false and false != false")
	}

	if node.Stringify() != "(false && false)" {
		t.Errorf("And not stringified correctly")
	}
}