package metrocard

import (
	"testing"
)

var m Money = 150

func TestString(t *testing.T) {
	if m.String() != "$1.50" {
		t.Error("Expected $1.50, got ", m.String())
	}
}

func TestSetTakesAString(t *testing.T) {
	m.Set("1.00")
	if m != 100 {
		t.Error("Expected 100, got ", m)
	}
}

func TestSetTakesAFloat(t *testing.T) {
	m.Set(1.00)
	if m != 100 {
		t.Error("Expected 100, got ", m)
	}
}
