package usecase

import (
	"context"

	"github.com/mohuishou/new-project/internal/domain"
)

type article struct {
	repo domain.IArticleRepo
	isTx bool
}

// NewArticleUsecase init
func NewArticleUsecase(repo domain.IArticleRepo) domain.IArticleUsecase {
	return &article{repo: repo}
}

func (u *article) GetArticle(ctx context.Context, id int) (*domain.Article, error) {
	// 这里可能有其他业务逻辑...
	return u.repo.GetArticle(ctx, u.repo.WithByID(uint(id)))
}

func (u *article) CreateArticle(ctx context.Context, article *domain.Article, tagIDs []uint) error {
	return u.repo.Tx(ctx, func(ctx context.Context, repo domain.IArticleRepo) error {
		err := repo.CreateArticle(ctx, article)
		if err != nil {
			return err
		}

		var ats []*domain.ArticleTag
		for _, tid := range tagIDs {
			ats = append(ats, &domain.ArticleTag{
				ArticleID: article.ID,
				TagID:     tid,
			})
		}
		return repo.CreateArticleTags(ctx, ats)
	})
}

type handler func(ctx context.Context, usecase domain.IArticleUsecase) error

func (u *article) tx(ctx context.Context, f handler) error {
	if u.isTx {
		return f(ctx, u)
	}

	return u.repo.Tx(ctx, func(ctx context.Context, repo domain.IArticleRepo) error {
		usecase := &article{
			repo: repo,
			isTx: true,
		}

		return f(ctx, usecase)
	})
}
