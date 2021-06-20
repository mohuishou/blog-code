package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	// 多跑几次来看结果
	for i := 0; i < 100000; i++ {
		run()
	}
}

func run() {
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go routine(i)
	}
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func routine(id int) {
	for i := 0; i < 2; i++ {
		value := counter
		value++
		counter = value
	}
	wg.Done()
}
