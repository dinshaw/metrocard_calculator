package metrocard

import (
	money "github.com/dinshaw/metrocard/money"
	"testing"
)

var value money.Money = 1000

// func TestNewCard(t *testing.T) {
// 	card := NewCard(value)
// 	if card.value != 1050 {
// 		t.Error("Expected card to have a value of 1050, but got", card.value)
// 	}
// }

// func TestExistingCard(t *testing.T) {
// 	card := ExistingCard(value)
// 	if card.value != value {
// 		t.Error("Expected card to have a value of 1000, but got", card.value)
// 	}
// }

// func TestRefill(t *testing.T) {
// 	card := ExistingCard(value)
// 	card.RefillCard(value)
// 	if card.value != 2050 {
// 		t.Error("Expected card to have a value of 2050, but got", card.value)
// 	}
// }

func TestValueWithBonus(t *testing.T) {
	value := valueWithBonus(1000)
	if value != 1050 {
		t.Error("Expected 1050, but got", value)
	}
}

// func TestCalculate(t *testing.T) {
// 	card := ExistingCard(value)
// 	value_to_add, _ := card.Calculate()
// 	if value_to_add != 905 {
// 		t.Error("Expected value_to_add to be 11, but got", value_to_add)
// 	}
// }
