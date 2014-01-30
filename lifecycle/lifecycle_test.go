package metrocard

import (
  "testing"
)

const expected = 14

func TestLifecycle(t *testing.T) {
  balances := Lifecycle(20.00, 1.50)
  if len(balances) != expected {
    t.Error("Expected %v, got ", expected, len(balances))
  }
}
