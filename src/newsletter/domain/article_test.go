package domain_test

import (
	"encoding/json"
	"io"
	"net/http"
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"pulzo/src/newsletter/view/dto"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFindArticleById(t *testing.T) {
	expectedSection := "deportes"
	articleRepository := mock.NewArticleDao()
	article := domain.FindArticleById("PP2891869", articleRepository)
	assert.Equal(t, article.Section, expectedSection)
}

func TestFindArticle(t *testing.T) {
	var findArticleJson dto.FindArticleDto
	urlJsonFile := "https://filesstaticpulzo.s3.us-west-2.amazonaws.com/pulzo-lite/posts/PP2865861A.json"

	response, err := http.Get(urlJsonFile)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(body, &findArticleJson)
	if err != nil {
		t.Error(err)
	}
}
