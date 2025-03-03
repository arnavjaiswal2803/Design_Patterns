package main

import "fmt"

type Lfu struct {
}

func (l *Lfu) evict() {
	fmt.Println("Evicting by lfu algo.")
}
