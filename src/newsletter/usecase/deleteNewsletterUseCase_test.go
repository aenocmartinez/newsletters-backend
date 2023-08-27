package usecase_test

import (
	"pulzo/src/newsletter/usecase"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestEliminarNewsletter(t *testing.T) {
	codeExpected := "200"
	casoUsoEliminar := usecase.DeleteNewsletterUseCase{}
	response := casoUsoEliminar.Execute(1)
	assert.Equal(t, response.Code, codeExpected)
}

func TestEliminarNewsNoExistente(t *testing.T) {
	codeExpected := "404"
	casoUsoEliminar := usecase.DeleteNewsletterUseCase{}
	response := casoUsoEliminar.Execute(12)
	assert.Equal(t, response.Code, codeExpected)
}
