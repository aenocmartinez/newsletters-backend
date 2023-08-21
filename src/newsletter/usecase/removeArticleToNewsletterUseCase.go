package usecase

import (
	"encoding/json"
	"pulzo/src/newsletter/domain"
	newsdto "pulzo/src/newsletter/view/dto"
	pulzoservices "pulzo/src/shared/infraestructure/pulzo_services"
	"pulzo/src/shared/view/dto"
)

type RemoveArticleToNewsletterUseCase struct{}

func (useCase *RemoveArticleToNewsletterUseCase) Execute(idNewsletter int64, posArticle int) (response dto.ResponseHttp) {

	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		response.Code = "404"
		response.Message = "newsletter not found"
		return response
	}

	newsletter.RemoveArticle(posArticle)

	content, _ := json.Marshal(newsletter.ArrayArticle())
	go pulzoservices.UploadFile(newsletter, "json", content)

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
