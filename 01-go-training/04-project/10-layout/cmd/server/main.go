package main

import "github.com/gin-gonic/gin"

func main() {
	s, err := NewServices()
	if err != nil {
		panic(err)
	}

	e := gin.Default()
	s.register(e)

	// 这里还有优雅中止的一些代码，就不贴了
}
