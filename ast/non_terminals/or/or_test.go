package or

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
		t.Errorf("true or true != true")
	}

	node.SetLeft(trueNode)
	node.SetRight(falseNode)
	
	if node.Interpret() != true {
		t.Errorf("true or false != true")
	}

	node.SetLeft(falseNode)
	node.SetRight(trueNode)
	
	if node.Interpret() != true {
		t.Errorf("false or true != true")
	}

	node.SetLeft(falseNode)
	node.SetRight(falseNode)
	
	if node.Interpret() != false {
		t.Errorf("false or false != false")
	}

	if node.Stringify() != "(false || false)" {
		t.Errorf("Or not stringified correctly")
	}
}