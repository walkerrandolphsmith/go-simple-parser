package terminal

import (
	"testing"
)

func TestNew(t *testing.T) {
	f := New(false)
	if f.Value != false {
		t.Errorf("False was not interpreted as false");
	}
}

func TestInterpret(t *testing.T) {
	f := New(false)
	if f.Interpret() != false {
		t.Errorf("False was not interpreted as false");
	}
}