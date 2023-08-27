package usecase_test

import (
	"pulzo/src/newsletter/usecase"
	"pulzo/src/newsletter/view/dto"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestListarNewslleter(t *testing.T) {
	numItemExpected := 3
	casoUsoListar := usecase.NewsletterListUseCase{}
	response := casoUsoListar.Execute()

	lista, _ := response.Data.([]dto.NewsletterDto)

	assert.Equal(t, len(lista), numItemExpected)
}
