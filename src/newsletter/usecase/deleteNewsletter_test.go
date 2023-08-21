package usecase_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDeleteNewsletter(t *testing.T) {

	var idNewsletter int64 = 8
	newsletterRepository := mock.NewNewsletterDao()
	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		t.Fatal("one (1) record expected")
	}
	newsletter.SetRepository(newsletterRepository)

	err := newsletter.Delete()
	assert.Equal(t, err, nil)
}
