package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg2 := wg
	wg2.Wait()
}
