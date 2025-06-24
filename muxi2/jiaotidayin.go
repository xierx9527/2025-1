package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	c := make(chan string)
	wg := sync.WaitGroup{}
	done := make(chan struct{})
	a := []string{}
	b := []string{}
	for i := 0; i < 26; i++ {
		a = append(a, string(65+i))
		b = append(b, strconv.Itoa(i+1))
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, i := range b {
			<-done
			c <- i
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, i := range a {
			c <- i
			done <- struct{}{}
			time.Sleep(5 * time.Millisecond)
		}
		close(done)
	}()
	go func() {
		wg.Wait()
		close(c)
	}()
	for val := range c {
		fmt.Println(val)
	}
}
