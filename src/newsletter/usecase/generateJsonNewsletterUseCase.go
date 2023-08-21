package usecase

import (
	"pulzo/src/newsletter/domain"
)

type GenerateJsonNewsletterUseCase struct{}

func (useCase *GenerateJsonNewsletterUseCase) Execute(newsletter domain.Newsletter) {

	// content, err := json.MarshalIndent(newsletter.JsonFileObject(), "", " ")
	// if err != nil {
	// 	sentry.CaptureException(err)
	// }

	// _, _, err = pulzoservices.UploadFile(newsletter.JsonFileName(), "json", content)
	// if err != nil {
	// 	sentry.CaptureException(err)
	// }

}
