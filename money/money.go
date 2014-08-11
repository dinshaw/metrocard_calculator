package metrocard

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type Money int64

// func (m *Money) String() string {
// 	return fmt.Sprintf("$%d.%02d", int64(*m)/100, int64(*m)%100)
// }

func (m *Money) Set(value interface{}) error {
	fmt.Printf("aoeu", reflect.TypeOf(value))
	switch value.(type) {
	case string:
		fmt.Print("String")
		float, err := strconv.ParseFloat(value.(string), 64)
		*m = Money(float * 100)
		// case float64:
		// 	*m = Money(value * 100)
		// case int:
		if err != nil {
			log.Fatal(err)
		}
		// 	*m = Money(value)
	}

	return nil
}
