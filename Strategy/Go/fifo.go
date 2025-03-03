package main

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict() {
	fmt.Println("evicting by Fifo algo")
}
