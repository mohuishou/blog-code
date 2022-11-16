package service

import (
	"context"

	v1 "github.com/mohuishou/new-project/api/product/app/v1"
	"github.com/mohuishou/new-project/internal/domain"
	"github.com/mohuishou/new-project/internal/pkg/copier"
)

// 确保实现了对应的接口
var _ v1.BlogServiceHTTPServer = &Artcile{}

// Artcile Artcile
type Artcile struct {
	usecase domain.IArticleUsecase
}

// NewArticleService 初始化方法
func NewArticleService(usecase domain.IArticleUsecase) *Artcile {
	return &Artcile{usecase: usecase}
}

// CreateArticle 创建一篇文章
func (a *Artcile) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (*v1.CreateArticleResp, error) {
	var article domain.Article
	err := copier.Copy(&article, req)
	if err != nil {
		return nil, err
	}

	err = a.usecase.CreateArticle(ctx, &article)
	return &v1.CreateArticleResp{}, err
}

// GetArticle 获取文章
func (a *Artcile) GetArticle(ctx context.Context, req *v1.GetArticleReq) (*v1.GetArticleResp, error) {
	return &v1.GetArticleResp{}, nil
}
