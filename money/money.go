package metrocard

import (
	"fmt"
	"log"
	"strconv"
)

type Money int64

func (m *Money) String() string {
	return fmt.Sprintf("$%d.%02d", int64(*m)/100, int64(*m)%100)
}

func (m *Money) Set(value string) error {

	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err)
	}
	*m = Money(float * 100)
	return nil
}
