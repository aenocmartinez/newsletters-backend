package mysql

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"pulzo/src/newsletter/view/dto"

	"github.com/getsentry/sentry-go"
)

type ArticleDao struct {
}

func NewArticleDao() *ArticleDao {
	return &ArticleDao{}
}

func (a *ArticleDao) FindArticleById(id string) (articleJson dto.ArticleDto) {
	var findArticleJson dto.FindArticleDto
	urlJsonFile := os.Getenv("PATH_ARTICLE") + "/" + id + ".json"

	response, err := http.Get(urlJsonFile)
	if err != nil {
		sentry.CaptureException(err)
		return articleJson
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		sentry.CaptureException(err)
		return articleJson
	}

	err = json.Unmarshal(body, &findArticleJson)
	if err != nil {
		sentry.CaptureException(err)
		return articleJson
	}

	articleJson.Id = findArticleJson.Id
	articleJson.Title = findArticleJson.Title.Main
	articleJson.Lead = findArticleJson.Lead.Main
	articleJson.Section = findArticleJson.Section.Main.Slug
	articleJson.Link = findArticleJson.Link.Main
	articleJson.Image = findArticleJson.Images.Types.Square.Medium
	articleJson.ImagenFA = findArticleJson.Images.Types.Horizontal.Big

	return articleJson
}
