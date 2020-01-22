package strats

import (
	cache "designpatterns/behavioural/strategy/cache"
	"fmt"
)

type LifoCachePolicy struct{}

func (lfu LifoCachePolicy) Replace(c *cache.Cache) {

	fmt.Println("Replaced Cache using LIFO")
	c.Data = c.Data[:len(c.Data)-1]
}
