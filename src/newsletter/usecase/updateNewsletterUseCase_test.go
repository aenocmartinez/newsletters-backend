package usecase_test

import (
	"pulzo/src/newsletter/usecase"
	"pulzo/src/newsletter/view/dto"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestActualizarNewsletter(t *testing.T) {
	expected := "200"
	newsletterDto := dto.NewsletterDto{
		Id:      1,
		Name:    "Entretenimientos",
		Section: "entretenimiento",
	}

	casoUsoActualizar := usecase.UpdateNewsletterUseCase{}
	response := casoUsoActualizar.Execute(newsletterDto)
	assert.Equal(t, response.Code, expected)
}

func TestActualizarNewsletterNoExistente(t *testing.T) {
	expected := "404"
	newsletterDto := dto.NewsletterDto{
		Id:      12,
		Name:    "Entretenimientos",
		Section: "entretenimiento",
	}

	casoUsoActualizar := usecase.UpdateNewsletterUseCase{}
	response := casoUsoActualizar.Execute(newsletterDto)
	assert.Equal(t, response.Code, expected)
}
