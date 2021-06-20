package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			LookupService("xx")
			RegisterService("xx", "11")
			wg.Done()
		}()
	}

	wg.Wait()
}

var service = map[string]string{}

// RegisterService RegisterService
func RegisterService(name, addr string) {
	service[name] = addr
}

// LookupService LookupService
func LookupService(name string) string {
	return service[name]
}
