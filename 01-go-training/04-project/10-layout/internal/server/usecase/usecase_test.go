package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_domain "github.com/mohuishou/new-project/internal/mock"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/mohuishou/new-project/internal/domain"
)

type articleHelper struct {
	repo *mock_domain.MockIArticleRepo
	article *article
}

func newArticleHelper(t *testing.T) *articleHelper {
	ctrl := gomock.NewController(t)
	h := &articleHelper{}
	h.repo = mock_domain.NewMockIArticleRepo(ctrl)
	h.article = NewArticleUsecase(h.repo).(*article)
	return h
}

func Test_article_CreateArticle(t *testing.T) {
	h := newArticleHelper(t)
	type args struct {
		ctx     context.Context
		article *domain.Article
		tagIDs  []uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.repo.EXPECT().
				Tx(gomock.Any(), gomock.Any()).
				Return(nil)
			err := h.article.CreateArticle(tt.args.ctx, tt.args.article, tt.args.tagIDs)
			if tt.wantErr {
				require.NotNil(t, err)
			}
			require.Nil(t, err)
		})
	}
}
