package main

import "fmt"

type Lru struct {
}

func (l *Lru) evict() {
	fmt.Println("evicting by lru algo")
}
