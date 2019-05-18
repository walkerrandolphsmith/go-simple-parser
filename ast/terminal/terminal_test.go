package terminal

import (
	"testing"
)

func TestNew(t *testing.T) {
	node := New(false)
	if node.Value != false {
		t.Errorf("False was not interpreted as false");
	}
}

func TestInterpret(t *testing.T) {
	node := New(false)
	if node.Interpret() != false {
		t.Errorf("False was not interpreted as false");
	}
}

func TestStringify(t *testing.T) {
	node := New(true)

	if node.Stringify() != "true" {
		t.Errorf("And is not stringified correctly")
	}
}