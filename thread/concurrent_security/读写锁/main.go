package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup
	rwlock sync.RWMutex
	x      int32
)

func write() {
	rwlock.Lock()
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	x += 1
	rwlock.Unlock()
	wg.Done()
}

func read() {
	rwlock.RLock()
	x += 1
	time.Sleep(1 * time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
