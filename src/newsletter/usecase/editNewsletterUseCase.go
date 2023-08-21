package usecase

import (
	"pulzo/src/newsletter/domain"
	"pulzo/src/shared/view/dto"

	newsdto "pulzo/src/newsletter/view/dto"
)

type EditNewsletterUseCase struct{}

func (useCase *EditNewsletterUseCase) Execute(idNewsletter int64) (response dto.ResponseHttp) {
	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		response.Code = "404"
		response.Message = "resource not found"
		return response
	}

	schedule := newsletter.Schedule()

	response.Code = "200"
	response.Data = newsdto.NewsletterDto{
		Id:             newsletter.Id(),
		Name:           newsletter.Name(),
		Section:        newsletter.Section(),
		State:          newsletter.State(),
		HtmlTemplate:   newsletter.HtmlTemplate(),
		Recipients:     newsletter.Recipients(),
		ScheduleStatus: schedule.Status(),
		ScheduleDate:   schedule.Date(),
		ScheduleHour:   schedule.Hour(),
	}

	return response
}
