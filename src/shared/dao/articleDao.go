package dao

import (
	"encoding/json"
	"os"
	"pulzo/src/shared/domain"
	"pulzo/src/shared/infraestructure/util"

	"github.com/getsentry/sentry-go"
)

type ArticleDao struct{}

func NewArticleDao() *ArticleDao {
	return &ArticleDao{}
}

func (a *ArticleDao) FindArticleById(id string) (article domain.Article) {
	urlJsonFile := os.Getenv("PATH_ARTICLE") + "/" + id + ".json"

	content, err := util.ReadFileFromURL(urlJsonFile)
	if err != nil {
		sentry.CaptureException(err)
		return article
	}

	err = json.Unmarshal(content, &article)
	if err != nil {
		sentry.CaptureException(err)
		return article
	}

	return article
}
