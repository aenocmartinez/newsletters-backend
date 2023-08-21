package usecase_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"testing"
)

func TestReadConfigNewsletter(t *testing.T) {
	t.Setenv("ROOT_JSON_FILE_PATH", "https://filesstaticpulzo.s3.us-west-2.amazonaws.com/pulzo-dev/jsons/admin")
	t.Setenv("APP_ENDPOINT_S3", "https://services.pulzo.com/pulzo-aws-stage/upload")
	t.Setenv("HTML_TEMPLATE_DEFAULT", "Template_Default")
	t.Setenv("APP_BUCKET_S3", "pulzo-dev/jsons/admin/")

	var idNewsletter int64 = 1
	newsletterRepository := mock.NewNewsletterDao()

	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		t.Fatal("expected ONE(1) record")
	}
}
