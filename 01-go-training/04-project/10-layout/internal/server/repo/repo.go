package repo

import (
	"context"

	"github.com/mohuishou/new-project/internal/domain"
	"gorm.io/gorm"
)

type article struct {
	db *gorm.DB
}

func (r *article) WithByID(id uint) domain.DBOption {
	panic("implement me")
}

func (r *article) WithByTitle(title string) domain.DBOption {
	panic("implement me")
}

func (r *article) GetArticleByTitle(ctx context.Context, title string) (*domain.Article, error) {
	panic("implement me")
}

func (r *article) GetArticleByID(ctx context.Context, id int) (*domain.Article, error) {
	panic("implement me")
}

// NewArticleRepo init
func NewArticleRepo(db *gorm.DB) domain.IArticleRepo {
	return &article{db: db}
}

func (r *article) GetArticle(ctx context.Context, opts ...domain.DBOption) (*domain.Article, error) {
	var a domain.Article
	db := r.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	if err := db.Find(&a); err != nil {
		// 这里返回业务错误码
	}
	return &a, nil
}

func (r *article) CreateArticle(ctx context.Context, article *domain.Article) error {
	if err := r.db.WithContext(ctx).Create(article); err != nil {
		// 这里返回业务错误码
	}
	return nil
}

func (r *article) Tx(ctx context.Context, f domain.ArticleRepoTxFunc) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := NewArticleRepo(tx)
		return f(ctx, repo)
	})
}

func (r *article) CreateArticleTags(ctx context.Context, ArticleTags []*domain.ArticleTag) error {
	return nil
}