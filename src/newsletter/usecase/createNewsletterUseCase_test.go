package usecase_test

import (
	"pulzo/src/newsletter/usecase"
	"pulzo/src/newsletter/view/dto"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCasoUsoCrearNewsletter(t *testing.T) {
	useCase := usecase.CreateNewsletterUseCase{}

	newsletterDto := dto.NewsletterDto{
		Name:       "Virales",
		Section:    "virales",
		Recipients: "11605",
	}

	response := useCase.Execute(newsletterDto)
	assert.Equal(t, response.Code, "201")
}

func TestCasoUsoCrearNewsletterYaExistente(t *testing.T) {
	useCase := usecase.CreateNewsletterUseCase{}

	newsletterDto := dto.NewsletterDto{
		Name:       "Entretenimiento",
		Section:    "entretenimiento",
		Recipients: "11605",
	}

	response := useCase.Execute(newsletterDto)
	assert.Equal(t, response.Code, "409")
}
