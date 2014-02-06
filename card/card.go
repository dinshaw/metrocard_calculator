package metrocard

import (
	// "fmt"
	"math"
)

var ride_cost int = 150
var bonus_percent int = 5
var min_increment int = 5

type card struct {
	value          int
	RidesRemaining int
}

func NewCard(value int) *card {
	card := new(card)
	card.value = valueWithBonus(value)
	return card
}

func (c *card) Value() int {
	return int(c.value)
}

func valueWithBonus(value int) int {
	bonus := int(math.Floor(float64(value) * float64(bonus_percent) * float64(0.01)))
	value_with_bonus := value + bonus
	return value_with_bonus
}

func ExistingCard(value int) *card {
	card := new(card)
	card.value = value
	return card
}

func (c *card) RefillCard(value int) bool {
	c.value += valueWithBonus(value)
	return true
}

func (c card) Calculate() (int, int) {
	initial_value := c.value
	value_to_add := 0
	for i := 0; math.Mod(float64(c.value), float64(ride_cost)) != 0; i += min_increment {
		c.value = initial_value

		// fmt.Printf("Starting Card value %v\n", c.value)
		// fmt.Printf("Adding %v, adjusted to %v\n", i, valueWithBonus(i))

		c.RefillCard(i)
		value_to_add = i

		// fmt.Printf("New Card value is %v\n", c.value)
		// fmt.Printf("Mod %v\n", math.Mod(float64(c.value), float64(ride_cost)))
		// fmt.Printf("=====================\n")
	}
	return value_to_add, (c.Value() / ride_cost)
}
