package main

import (
	"fmt"
	"sync"
)

var x int32
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 500; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
