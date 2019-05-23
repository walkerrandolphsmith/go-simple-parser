package lexer

import (
	"testing"
)

func TestNew(t *testing.T) {
	input := "(true or false) and not false"
	lexer := New(input)
	if lexer.tokens == nil {
		t.Errorf("new lexer does not have input")
	}
	if lexer.index != -1 {
		t.Errorf("new lexer has incorrect index")
	}
}

func TestNextSymbol(t *testing.T) {
	lexer := New("not false or (true and false)")
	token := lexer.NextSymbol()
	if token != NOT {
		t.Errorf("NOT")
	}

	token = lexer.NextSymbol()
	if token != FALSE {
		t.Errorf("FALSE")
	}

	token = lexer.NextSymbol()
	if token != OR {
		t.Errorf("OR")
	}

	token = lexer.NextSymbol()
	if token != LEFT {
		t.Errorf("LEFT")
	}

	token = lexer.NextSymbol()
	if token != TRUE {
		t.Errorf("TRUE")
	}

	token = lexer.NextSymbol()
	if token != AND {
		t.Errorf("AND")
	}

	token = lexer.NextSymbol()
	if token != FALSE {
		t.Errorf("FALSE")
	}

	token = lexer.NextSymbol()
	if token != RIGHT {
		t.Errorf("RIGHT")
	}

	token = lexer.NextSymbol()
	if token != INVALID {
		t.Errorf("INVALID")
	}
}