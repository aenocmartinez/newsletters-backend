package usecase

import "pulzo/src/shared/domain"

type FindArticleByUrlUseCase struct{}

func (useCase *FindArticleByUrlUseCase) Execute(url string) (article domain.Article, err error) {
	return domain.FindArticleByURL(url, articleRepository)
}
