package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	v1 "github.com/mohuishou/new-project/api/product/app/v1"
	"github.com/mohuishou/new-project/config/initializer"
	"github.com/mohuishou/new-project/internal/server"
	"github.com/mohuishou/new-project/internal/server/service"
)

var set = wire.NewSet(
	// domains
	server.Set,

	// common
	initializer.Set,
)

type services struct {
	article *service.Artcile
}

func (s *services) register(r gin.IRouter) {
	v1.RegisterBlogServiceHTTPServer(r, s.article)
}
