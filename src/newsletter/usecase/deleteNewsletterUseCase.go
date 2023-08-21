package usecase

import (
	"pulzo/src/newsletter/domain"
	"pulzo/src/shared/view/dto"
)

type DeleteNewsletterUseCase struct{}

func (useCase *DeleteNewsletterUseCase) Execute(idNewsletter int64) (response dto.ResponseHttp) {
	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		response.Code = "404"
		response.Message = "resource not found"
		return response
	}

	newsletter.SetRepository(newsletterRepository)

	err := newsletter.Delete()
	if err != nil {
		response.Code = "500"
		response.Message = err.Error()
		return response
	}

	response.Code = "200"
	response.Message = "success"
	return response
}
