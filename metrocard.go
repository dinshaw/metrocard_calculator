package main

import (
	"flag"
	"fmt"
	"github.com/dinshaw/metrocard/card"
	"log"
	"strconv"
)

type money int

func (m *money) String() string {
	// return fmt.Sprintf("$%.2d", *m)
	return "aoeu"
}

func (m *money) Set(value string) error {
	float, err := strconv.ParseFloat(value, 1)
	if err != nil {
		log.Fatal(err)
	}
	*m = money(float * 100)
	return nil
}

var moneyFlag money
var value_to_add money

func init() {
	flag.Var(&moneyFlag, "existing", "How much is currently on your Metrocard")
}

func main() {
	flag.Parse()
	c := metrocard.ExistingCard(int(moneyFlag))
	value_to_add, rides_remaining := c.Calculate()
	fmt.Printf("If you add %v you will have a zero balance after %v rides.\n", money(value_to_add), rides_remaining)

}
