package usecase

import (
	"pulzo/src/newsletter/domain"
	newsdto "pulzo/src/newsletter/view/dto"
	"pulzo/src/shared/view/dto"
)

type NewsletterListUseCase struct{}

func (useCase *NewsletterListUseCase) Execute() (response dto.ResponseHttp) {
	list := domain.NewsletterList(newsletterRepository)
	if len(list) == 0 {
		list = []newsdto.NewsletterDto{}
	}
	response.Code = "200"
	response.Data = list

	return response
}
