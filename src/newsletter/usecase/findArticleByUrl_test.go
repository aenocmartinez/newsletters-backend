package usecase_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFindArticleByUrl(t *testing.T) {
	var expectedId string = "PP2891869"
	articleRepository := mock.NewArticleDao()

	var url string = "/deportes/luis-fernando-muriel-estara-despedida-sebastian-viera-metropolitano-PP2891869"
	idArticle := domain.ExtractIdArticleFromURL(url)

	if idArticle == "" {
		t.Error("article not found")
	}

	articleDto := domain.FindArticleById(idArticle, articleRepository)
	assert.Equal(t, articleDto.Id, expectedId)
}
