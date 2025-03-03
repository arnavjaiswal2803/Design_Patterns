package main

func main() {
	lru, fifo, lfu := Lru{}, Fifo{}, Lfu{}

	cache := InitCache(&lru)

	cache.add()
	cache.add()

	cache.add()

	cache.setEvictionAlgo(&fifo)
	cache.add()

	cache.setEvictionAlgo(&lfu)
	cache.add()
}
