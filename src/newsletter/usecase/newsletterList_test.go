package usecase_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNewsletterList(t *testing.T) {
	axpectedRecords := 10
	newsletterRepository := mock.NewNewsletterDao()
	list := domain.NewsletterList(newsletterRepository)
	assert.Equal(t, len(list), axpectedRecords)
}
