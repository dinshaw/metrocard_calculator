package metrocard

import (
  // "fmt"
)

func Lifecycle(value, cost float64) ([]float64) {
  balances := make([]float64, (int(value/cost) + 1))
  for i := 0; value >= 0; i++ {
    balances[i] = value
    value -= cost
  }
  return balances
}
