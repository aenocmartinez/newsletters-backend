package usecase

import (
	"pulzo/src/newsletter/domain"
	"pulzo/src/shared/view/dto"

	newsdto "pulzo/src/newsletter/view/dto"
)

type SaveSettingNewsletterUseCase struct{}

func (useCase *SaveSettingNewsletterUseCase) Execute(scheduleDto newsdto.ScheduleNewsletterDto) (response dto.ResponseHttp) {
	newsletter := domain.FindNewsletterById(scheduleDto.Id, newsletterRepository)
	if !newsletter.Exists() {
		response.Code = "404"
		response.Message = "newsletter not found"
		return response
	}

	newsletter.SetRepository(newsletterRepository)
	newsletter.ReadFileJson()

	schedule := newsletter.Schedule()
	if schedule.WasExecuted() {
		schedule.SetCampaignId("")
	}

	schedule.SetDate(scheduleDto.Date)
	schedule.SetHour(scheduleDto.Hour)
	schedule.SetSubject(scheduleDto.Subject)
	schedule.SetFrom(scheduleDto.From)
	schedule.SetMessage(newsletter.GenerateHtml())

	newsletter.SetSchedule(schedule)

	if newsletter.IsMailchimp() {
		sendMailchimp := SendMailchimpUseCase{}
		go sendMailchimp.Execute(newsletter)
	}

	response.Code = "200"
	response.Message = "success"
	return response
}
