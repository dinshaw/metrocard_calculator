package metrocard

import (
	"fmt"
	money "github.com/dinshaw/metrocard/money"
	"math"
)

var ride_cost money.Money = 150
var bonus_percent int = 5
var min_increment int = 5

type card struct {
	value          money.Money
	RidesRemaining int
}

func NewCard(value money.Money) *card {
	card := new(card)
	card.value = valueWithBonus(value)
	return card
}

func ExistingCard(value money.Money) *card {
	card := new(card)
	card.value = value
	return card
}

func (c *card) RefillCard(value money.Money) bool {
	c.value += valueWithBonus(value)
	return true
}

func (c card) Calculate() (money.Money, int) {
	initial_value := c.value
	value_to_add := money.Money(0)
	for i := 0; math.Mod(float64(c.value), float64(ride_cost)) != 0; i += min_increment {
		c.value = initial_value

		fmt.Printf("Starting Card value %v\n", c.value)
		fmt.Printf("Adding %v, adjusted to %v\n", i, valueWithBonus(money.Money(i)))

		c.RefillCard(money.Money(i))
		value_to_add = money.Money(i)

		fmt.Printf("New Card value is %v\n", c.value)
		fmt.Printf("Mod %v\n", math.Mod(float64(c.value), float64(ride_cost)))
		fmt.Printf("=====================\n")
	}
	return value_to_add, int(c.value / ride_cost)
}

func valueWithBonus(value money.Money) money.Money {
	value_as_float := float64(value) * 0.01
	bonus_percentage := float64(bonus_percent) * 0.01
	bonus_amount := value_as_float * bonus_percentage
	value_with_bonus := value + bonus_amount
	return value_with_bonus
}
