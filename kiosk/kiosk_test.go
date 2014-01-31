package metrocard

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	card := NewCard(10.00)
	if card.Value != 10.50 {
		t.Error("Expected 10.50, got ", card.Value)
	}
}

func TestRefillCard(t *testing.T) {
	card := NewCard(10.00)
	RefillCard(&card, 10.00)
	if card.Value != 21.00 {
		t.Error("Expected 21.00, got ", card.Value)
	}
}
