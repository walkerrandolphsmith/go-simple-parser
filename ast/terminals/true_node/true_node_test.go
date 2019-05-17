package true_node

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	node := New()
	if node.Interpret() != true {
		t.Errorf("True was not interpreted as true");
	}
}