package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int
var mu sync.Mutex

func main() {
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go routine(i)
	}
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func routine(id int) {
	for i := 0; i < 2; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
	wg.Done()
}
