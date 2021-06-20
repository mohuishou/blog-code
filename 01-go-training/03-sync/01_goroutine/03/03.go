package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func setup() {
	// 这里面有一些初始化的操作
}

func main() {
	setup()

	// for debug
	go pprof()

	// 主服务
	server()
}

func server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// 主服务
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Panicf("http server err: %+v", err)
		return
	}
}

func pprof() {
	// 辅助服务，监听了其他端口，这里是 pprof 服务，用于 debug
	http.ListenAndServe(":8081", nil)
}
