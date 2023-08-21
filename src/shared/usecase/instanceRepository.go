package usecase

import (
	"pulzo/src/shared/dao"
	"pulzo/src/shared/domain"
)

var articleRepository domain.ArticleRepository = dao.NewArticleDao()
