package main

import (
	"fmt"
	"time"
)

func main() {
	// буферизованным каналом можно ограничить кол-во
	// паралельно выполняемых корутин
	lock := make(chan bool, 2)
	for i := 1; i <= 6; i++ {
		go worker(i, lock)
	}
	time.Sleep(5 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", id)
	<-lock
}
