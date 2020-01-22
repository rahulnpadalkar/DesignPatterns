package strats

import (
	cache "designpatterns/behavioural/strategy/cache"
	"fmt"
)

type FifoCachePolicy struct{}

func (fifo FifoCachePolicy) Replace(c *cache.Cache) {

	fmt.Println("Replaced Cache using Fifo")
	c.Data = c.Data[2:]
}
