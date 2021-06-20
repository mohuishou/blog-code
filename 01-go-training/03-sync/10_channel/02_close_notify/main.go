package main

import (
	"fmt"
	"time"
)

func run(stop <-chan struct{}, done chan<- struct{}) {
	// 每一秒打印一次 hello
	for {
		select {
		case <-stop:
			fmt.Println("stop...")
			done <- struct{}{}
			return
		case <-time.After(time.Second):
			fmt.Println("hello")
		}
	}
}

func main() {
	// 一对多
	stop := make(chan struct{})
	// 多对一
	done := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go run(stop, done)
	}

	// 5s 后退出
	time.Sleep(5 * time.Second)
	close(stop)

	for i := 0; i < 10; i++ {
		<-done
	}
}
