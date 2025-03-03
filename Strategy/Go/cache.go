package main

type Cache struct {
	evictionAlgo EvictionAlgo // This is the strategy
	capacity     int
	maxCapacity  int
}

func InitCache(evictionAlgo EvictionAlgo) *Cache {
	return &Cache{
		evictionAlgo: evictionAlgo,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) add() {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
}

func (c *Cache) evict() {
	c.evictionAlgo.evict()
	c.capacity--
}

func (c *Cache) setEvictionAlgo(evictionAlgo EvictionAlgo) {
	c.evictionAlgo = evictionAlgo
}
