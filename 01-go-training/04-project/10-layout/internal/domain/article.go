package domain

import (
	"context"

	"gorm.io/gorm"
)

// Article 文章
type Article struct {
	Model // 基础结构体，包含 id, created_at, deleted_at, updated_at

	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `json:"tags" gorm:"many2many:article_tags"`
}

// IArticleUsecase IArticleUsecase
type IArticleUsecase interface {
	GetArticle(ctx context.Context, id int) (*Article, error)
	CreateArticle(ctx context.Context, article *Article, tagIDs []uint) error
}

// ArticleRepoTxFunc 事务方法
type ArticleRepoTxFunc = func(ctx context.Context, repo IArticleRepo) error
type DBOption func(*gorm.DB) *gorm.DB

//go:generate mockgen -destination=../mock/article.go . IArticleRepo
// IArticleRepo IArticleRepo
type IArticleRepo interface {
	Tx(ctx context.Context, f ArticleRepoTxFunc) error
	WithByID(id uint) DBOption
	WithByTitle(title string) DBOption
	GetArticle(ctx context.Context, opts ...DBOption) (*Article, error)
	GetArticleByTitle(ctx context.Context, title string) (*Article, error)
	GetArticleByID(ctx context.Context, id int) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
	CreateArticleTags(ctx context.Context, ArticleTags []*ArticleTag) error
}
