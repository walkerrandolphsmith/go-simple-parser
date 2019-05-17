package false_node

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	node := New()
	if node.Interpret() != false {
		t.Errorf("False was not interpreted as false");
	}
}