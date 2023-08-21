package domain_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreateDefaultNewsletter(t *testing.T) {
	newsletterRepository := mock.NewNewsletterDao()

	t.Setenv("HTML_TEMPLATE_DEFAULT", "Template_Default")

	newsletter := domain.NewNewsletter("Test 1")
	newsletter.SetSection("Section 1")
	newsletter.SetRepository(newsletterRepository)

	err := newsletter.Create()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateCustomNewsletter(t *testing.T) {
	newsletterRepository := mock.NewNewsletterDao()

	t.Setenv("HTML_TEMPLATE_DEFAULT", "Template_Default")

	newsletter := domain.NewNewsletter("Test 1")
	newsletter.SetSection("Section 1")
	newsletter.SetHtmlTemplate("Custom_Template")
	newsletter.SetState("active")
	newsletter.SetRepository(newsletterRepository)

	err := newsletter.Create()
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateNewsletter(t *testing.T) {
	// var expectedHour string = "10:25:00"
	t.Setenv("HTML_TEMPLATE_DEFAULT", "Template_Default")

	newsletterRepository := mock.NewNewsletterDao()
	newsletter := newsletterRepository.FindNewsletterById(5)
	newsletter.SetRepository(newsletterRepository)
	// newsletter.SetHour("10:25:00")

	err := newsletter.Update()
	if err != nil {
		t.Error(err)
	}

	// assert.Equal(t, newsletter.Hour(), expectedHour)
}

func TestDeleteNewsletter(t *testing.T) {
	newsletterRepository := mock.NewNewsletterDao()
	newsletter := newsletterRepository.FindNewsletterById(7)
	newsletter.SetRepository(newsletterRepository)
	err := newsletter.Delete()
	assert.Equal(t, err, nil)
}

func TestNewsletterList(t *testing.T) {
	expectedRecords := 10
	newsletterRepository := mock.NewNewsletterDao()
	arrayNewsletter := domain.NewsletterList(newsletterRepository)
	assert.Equal(t, len(arrayNewsletter), expectedRecords)
}

func TestFindNewsletterById(t *testing.T) {
	var expectedSection string = "Section_1"
	newsletterRepository := mock.NewNewsletterDao()
	newsletter := domain.FindNewsletterById(1, newsletterRepository)
	assert.Equal(t, newsletter.Section(), expectedSection)
}

func TestFindNewsletterByName(t *testing.T) {
	var expectedSection string = "Section_5"
	newsletterRepository := mock.NewNewsletterDao()
	newsletter := domain.FindNewsletterByName("Name newsletter 5", newsletterRepository)
	assert.Equal(t, newsletter.Section(), expectedSection)
}

func TestGetNewsletterJsonFile(t *testing.T) {
	t.Setenv("ROOT_JSON_FILE_PATH", "https://filesstaticpulzo.s3.us-west-2.amazonaws.com/pulzo-dev/jsons/admin")
	t.Setenv("APP_ENDPOINT_S3", "https://services.pulzo.com/pulzo-aws-stage/upload")
	t.Setenv("HTML_TEMPLATE_DEFAULT", "Template_Default")
	t.Setenv("APP_BUCKET_S3", "pulzo-dev/jsons/admin/")

	newsletterRepository := mock.NewNewsletterDao()
	newsletter := domain.NewNewsletter("Name newsletter test")
	newsletter.SetRepository(newsletterRepository)
}
