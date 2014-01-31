package metrocard

import (
	"github.com/dinshaw/metrocard/card"
)

var bonus_percent float64 = 0.05

func NewCard(value float64) metrocard.Card {
	card := metrocard.Card{Value: (value + (value * bonus_percent))}
	return card
}

func RefillCard(card *metrocard.Card, value float64) bool {
	card.Value += (value + (value * bonus_percent))
	return true
}
