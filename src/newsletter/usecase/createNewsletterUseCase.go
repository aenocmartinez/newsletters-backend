package usecase

import (
	"pulzo/src/newsletter/domain"
	newsdto "pulzo/src/newsletter/view/dto"
	"pulzo/src/shared/view/dto"
)

type CreateNewsletterUseCase struct{}

func (useCase *CreateNewsletterUseCase) Execute(newsletterDto newsdto.NewsletterDto) (response dto.ResponseHttp) {

	newsletter := domain.FindNewsletterByName(newsletterDto.Name, newsletterRepository)
	if newsletter.Exists() {
		response.Code = "409"
		response.Message = "the resource already exists"
		return response
	}

	newsletter = *domain.NewNewsletter(newsletterDto.Name)
	newsletter.SetRepository(newsletterRepository)
	newsletter.SetSection(newsletterDto.Section)
	newsletter.SetHtmlTemplate(newsletterDto.HtmlTemplate)
	newsletter.SetRecipients(newsletterDto.Recipients)

	err := newsletter.Create()
	if err != nil {
		response.Code = "500"
		response.Message = err.Error()
		return response
	}

	response.Code = "201"
	response.Message = "Susccessful"

	return response
}
