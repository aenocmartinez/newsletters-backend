package usecase

import (
	"pulzo/src/newsletter/domain"
	newsdto "pulzo/src/newsletter/view/dto"
	"pulzo/src/shared/view/dto"
)

type ScheduleNewslletterUseCase struct{}

func (useCase *ScheduleNewslletterUseCase) Execute(idNewsletter int64) (response dto.ResponseHttp) {
	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		response.Code = "404"
		response.Message = "resource not found"
		return response
	}

	newsletter.SetRepository(newsletterRepository)
	newsletter.ReadFileJson()

	if newsletter.ScheduleWasExecuted() {
		newsletter.DeleteSchedule()
	}

	response.Code = "200"
	response.Data = newsdto.FormSettingNewsletterDto{
		IdNewsletter: newsletter.Id(),
		Date:         newsletter.ScheduleDate(),
		Hour:         newsletter.ScheduleHour(),
		Content:      newsletter.Setting(),
		Status:       newsletter.ScheduleStatus(),
		Subject:      newsletter.ScheduleSubject(),
	}

	return response
}
