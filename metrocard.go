package main

import (
	"fmt"
)
var ride_cost float64 = 1.50

func main() {
	card := metrocard.ExistingCard(starting_balance)
  card.Calculate
	fmt.Printf(
    "If you add $%v you will have a zero balance after %v rides.\n"
    card.AmountToAdd,
    card.RidesRemaining
  )
}
