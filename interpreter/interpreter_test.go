package interpreter

import (
	"testing"
)

func TestEval(t *testing.T) {
	result := Eval("true and true")
	if result != true {
		t.Errorf("true and true")
	}

	result = Eval("true and false")
	if result != false {
		t.Errorf("true and false")
	}

	result = Eval("(true and false) or true")
	if result != true {
		t.Errorf("(true and false) or true")
	}

	result = Eval("(true and (true or false)) or false")
	if result != true {
		t.Errorf("(true and (true or false)) or false")
	}

	result = Eval("(false or true)")
	if result != true {
		t.Errorf("(false or true)")
	}
}