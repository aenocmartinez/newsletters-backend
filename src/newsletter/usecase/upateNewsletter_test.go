package usecase_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"pulzo/src/newsletter/usecase"
	"pulzo/src/shared/infraestructure/util"
	"testing"
	"time"
)

func TestUpdateNewsletter(t *testing.T) {
	t.Setenv("ROOT_JSON_FILE_PATH", "https://filesstaticpulzo.s3.us-west-2.amazonaws.com/pulzo-dev/jsons/admin")
	t.Setenv("APP_ENDPOINT_S3", "https://services.pulzo.com/pulzo-aws-stage/upload")
	t.Setenv("HTML_TEMPLATE_DEFAULT", "Template_Default")
	t.Setenv("APP_BUCKET_S3", "pulzo-dev/jsons/admin/")

	var idNewsletter int64 = 10
	newsletterRepository := mock.NewNewsletterDao()

	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		t.Fatal("one (1) record expected")
	}

	oldName := newsletter.Name()
	date := time.Now().Format(util.YYYY_MM_DD)

	newsletter.SetRepository(newsletterRepository)
	newsletter.SetName("Name newsletter updated " + date)
	newsletter.SetHtmlTemplate("")

	err := newsletter.Update()
	if err != nil {
		t.Fatal(err)
	}

	if oldName == newsletter.Name() {
		return
	}

	generatedJson := usecase.GenerateJsonNewsletterUseCase{}
	generatedJson.Execute(newsletter)
}
