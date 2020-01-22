package main

import (
	cache "designpatterns/behavioural/strategy/cache"
	strats "designpatterns/behavioural/strategy/strats"
	"fmt"
)

func main() {

	c := cache.InitCache()

	c.Add("1")
	c.Add("2")
	c.Add("3")
	c.Add("4")
	c.Add("5")
	c.Add("6")

	//Using Fifo
	fifo := strats.FifoCachePolicy{}
	c.SetPolicy(fifo)
	c.Policy.Replace(c)
	fmt.Println(c.Data)

	//Using Lifo
	lifo := strats.LifoCachePolicy{}
	c.SetPolicy(lifo)
	c.Policy.Replace(c)
	fmt.Println(c.Data)

	//Using LRU
	lru := strats.LruCachePolicy{}
	c.SetPolicy(lru)
	c.Policy.Replace(c)
	fmt.Println(c.Data)

}
