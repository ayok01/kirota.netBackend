package interactor

import (
	"context"

	"github.com/kirota401e/kirota.netBackend/entity/model"
)

type Article struct {
}

type IArticleInteractor interface {
	Get(ctx context.Context, articleId int) (*model.Article, error)
}