package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/gin-gonic/gin"
)

var start time.Time

func init() {
	start = time.Now()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	go server()
	hystrix.ConfigureCommand("test", hystrix.CommandConfig{
		// 执行 command 的超时时间
		Timeout: 0,

		// 最大并发量
		MaxConcurrentRequests: 100,

		// 一个统计窗口 10 秒内请求数量
		// 达到这个请求数量后才去判断是否要开启熔断
		RequestVolumeThreshold: 10,

		// 熔断器被打开后
		// SleepWindow 的时间就是控制过多久后去尝试服务是否可用了
		SleepWindow: 500,

		// 错误百分比
		// 请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个百分比后就会启动熔断
		ErrorPercentThreshold: 20,
	})

	for i := 0; i < 20; i++ {
		_ = hystrix.Do("test", func() error {
			resp, _ := resty.New().R().Get("http://localhost:8080/ping")
			if resp.IsError() {
				return fmt.Errorf("err code: %s", resp.Status())
			}
			return nil
		}, func(err error) error {
			fmt.Println("fallback err: ", err)
			return err
		})
		time.Sleep(100 * time.Millisecond)
	}

}

func server() {
	e := gin.Default()
	e.GET("/ping", func(ctx *gin.Context) {
		if time.Since(start) < 201*time.Millisecond {
			ctx.String(http.StatusInternalServerError, "pong")
			return
		}
		ctx.String(http.StatusOK, "pong")
	})
	e.Run(":8080")
}
