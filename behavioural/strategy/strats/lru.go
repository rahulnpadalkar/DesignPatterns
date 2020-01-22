package strats

import (
	cache "designpatterns/behavioural/strategy/cache"
	"fmt"
	"math/rand"
)

type LruCachePolicy struct{}

func (lru LruCachePolicy) Replace(cache *cache.Cache) {

	fmt.Println("Replaced Cache using LRU")
	i := rand.Intn(3)
	copy(cache.Data[i:], cache.Data[i+1:])
	cache.Data[len(cache.Data)-1] = ""
	cache.Data = cache.Data[:len(cache.Data)-1]
}
