package server

import (
	"github.com/google/wire"
	"github.com/mohuishou/new-project/internal/server/repo"
	"github.com/mohuishou/new-project/internal/server/service"
	"github.com/mohuishou/new-project/internal/server/usecase"
)

// Set for di
var Set = wire.NewSet(
	service.NewArticleService,
	usecase.NewArticleUsecase,
	repo.NewArticleRepo,
)
