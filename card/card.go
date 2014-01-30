package metrocard

type Card struct {
  Value float64
}

func (c *Card) Refill(value float64) bool {
  c.Value += value
  return true
}
