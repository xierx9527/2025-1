package main

import (
	"fmt"
	"sync"
)

var (
	res = 0
	mu  sync.Mutex
)

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		res += 1
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go add(&wg)
	}
	wg.Wait()
	fmt.Println(res)
}
