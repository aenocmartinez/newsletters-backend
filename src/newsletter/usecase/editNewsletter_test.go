package usecase_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestEditNewsletter(t *testing.T) {
	var expectedResponse bool = true
	var idNewsletter int64 = 8
	newsletterRepository := mock.NewNewsletterDao()
	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	assert.Equal(t, newsletter.Exists(), expectedResponse)
}
