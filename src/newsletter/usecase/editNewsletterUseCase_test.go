package usecase_test

import (
	"pulzo/src/newsletter/usecase"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestEditarNewsletter(t *testing.T) {
	codeExpected := "200"
	casoUsoEditar := usecase.EditNewsletterUseCase{}
	response := casoUsoEditar.Execute(2)
	assert.Equal(t, response.Code, codeExpected)
}

func TestEditarNewsletterNoExistente(t *testing.T) {
	codeExpected := "404"
	casoUsoEditar := usecase.EditNewsletterUseCase{}
	response := casoUsoEditar.Execute(12)
	assert.Equal(t, response.Code, codeExpected)
}
