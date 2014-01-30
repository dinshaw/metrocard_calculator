package metrocard

import (
  "testing"
)
var value float64 = 20.00

func TestRefill(t *testing.T) {
  card := Card{Value: value}
  card.Refill(value)
  if card.Value != (value * 2) {
    t.Error("Expected card to have a value of", (value * 2), "got", card.Value)
  }
}

