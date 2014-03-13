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
	switch t := value.(type) {
	case string:
		float, err := strconv.ParseFloat(value, 64)
		*m = Money(float * 100)
		// case float64:
		// 	*m = Money(value * 100)
		// case int:
		// 	*m = Money(value)
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
