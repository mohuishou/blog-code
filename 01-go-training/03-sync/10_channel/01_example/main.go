package main

import (
	"fmt"
)

func read(c <-chan int) {
	fmt.Println("read:", <-c)
}

func write(c chan<- int) {
	c <- 0
}

func main() {
	c := make(chan int)
	go read(c)
	write(c)
}
