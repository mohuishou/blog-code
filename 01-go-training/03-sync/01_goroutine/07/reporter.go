package main

import (
	"fmt"
	"sync"
	"time"
)

// Reporter 埋点服务上报
type Reporter struct {
	worker   int
	messages chan string
	wg       sync.WaitGroup
	closed   bool
}

// NewReporter NewReporter
func NewReporter(worker, buffer int) *Reporter {
	return &Reporter{worker: worker, messages: make(chan string, buffer)}
}

func (r *Reporter) run(stop <-chan struct{}) {
	go func() {
		<-stop
		r.shutdown()
	}()

	for i := 0; i < r.worker; i++ {
		r.wg.Add(1)
		go func() {
			for msg := range r.messages {
				time.Sleep(5 * time.Second)
				fmt.Printf("report: %s\n", msg)
			}
			r.wg.Done()
		}()
	}
	r.wg.Wait()
}

func (r *Reporter) shutdown() {
	r.closed = true
	// 注意，这个一定要在主服务结束之后再执行，避免关闭 channel 还有其他地方在啊写入
	close(r.messages)
}

// 模拟耗时
func (r *Reporter) report(data string) {
	if r.closed {
		return
	}
	r.messages <- data
}
