package usecase

import (
	"pulzo/src/newsletter/domain"
	newsdto "pulzo/src/newsletter/view/dto"
	"pulzo/src/shared/view/dto"
)

type UpdateNewsletterUseCase struct{}

func (useCase *UpdateNewsletterUseCase) Execute(newsletterDto newsdto.NewsletterDto) (response dto.ResponseHttp) {

	newsletter := domain.FindNewsletterById(newsletterDto.Id, newsletterRepository)
	if !newsletter.Exists() {
		response.Code = "404"
		response.Message = "resource not found"
		return response
	}

	oldName := newsletter.Name()

	newsletter.SetRepository(newsletterRepository)
	newsletter.SetName(newsletterDto.Name)
	newsletter.SetSection(newsletterDto.Section)
	newsletter.SetHtmlTemplate(newsletterDto.HtmlTemplate)
	newsletter.SetRecipients(newsletterDto.Recipients)
	newsletter.SetState(newsletterDto.State)

	err := newsletter.Update()
	if err != nil {
		response.Code = "500"
		response.Message = err.Error()
		return response
	}

	if oldName == newsletter.Name() {
		response.Code = "200"
		response.Message = "seccess"
		return response
	}

	response.Code = "200"
	response.Message = "seccess"

	return response
}
