package metrocard

import (
	"testing"
)

var value int = 1000

func TestNewCard(t *testing.T) {
	card := NewCard(value)
	if card.Value() != 1050 {
		t.Error("Expected card to have a value of 1050, but got", card.Value())
	}
}

func TestExistingCard(t *testing.T) {
	card := ExistingCard(value)
	if card.Value() != value {
		t.Error("Expected card to have a value of 1000, but got", card.Value())
	}
}

func TestRefill(t *testing.T) {
	card := ExistingCard(value)
	card.RefillCard(value)
	if card.Value() != 2050 {
		t.Error("Expected card to have a value of 2050, but got", card.Value())
	}
}

func TestCalculate(t *testing.T) {
	card := ExistingCard(value)
	value_to_add, _ := card.Calculate()
	if value_to_add != 905 {
		t.Error("Expected value_to_add to be 11, but got", value_to_add)
	}
}
