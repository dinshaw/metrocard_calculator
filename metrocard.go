package main

import (
	"flag"
	"fmt"
	"github.com/dinshaw/metrocard/card"
	money "github.com/dinshaw/metrocard/money"
)

var existing money.Money

func init() {
	flag.Var(&existing, "existing", "How much is currently on your Metrocard")
}

func main() {
	flag.Parse()
	c := metrocard.ExistingCard(existing)
	value_to_add, rides_remaining := c.Calculate()
	fmt.Printf("If you add %6.2d you will have a zero balance after %v rides.\n", value_to_add, rides_remaining)
}
