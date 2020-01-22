package cache

type Cache struct {
	Policy CacheReplacemnetAlgo
	Data   []string
}

func (c *Cache) Add(s string) {
	c.Data = append(c.Data, s)
}

func InitCache() *Cache {

	return &Cache{
		Policy: nil,
		Data:   []string{""},
	}
}

func (c *Cache) SetPolicy(policy CacheReplacemnetAlgo) {
	c.Policy = policy
}
