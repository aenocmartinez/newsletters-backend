package usecase

import (
	"encoding/json"
	"pulzo/src/newsletter/domain"
	newsdto "pulzo/src/newsletter/view/dto"

	pulzoservices "pulzo/src/shared/infraestructure/pulzo_services"
	"pulzo/src/shared/view/dto"
)

type AddArticleToNewsletterUseCase struct{}

func (useCase *AddArticleToNewsletterUseCase) Execute(url string, idNewsletter int64) (response dto.ResponseHttp) {

	newsletter := domain.FindNewsletterById(idNewsletter, newsletterRepository)
	if !newsletter.Exists() {
		response.Code = "404"
		response.Message = "newsletter not found"
		return response
	}

	idArticle := domain.ExtractIdArticleFromURL(url)
	if len(idArticle) == 0 {
		response.Code = "500"
		response.Message = "id article not found in the url"
		return response
	}

	articleDto := domain.FindArticleById(idArticle, articleRepository)
	if len(articleDto.Id) == 0 {
		response.Code = "404"
		response.Message = "article not found"
		return response
	}

	newsletter.SetRepository(newsletterRepository)
	newsletter.AddArticle(articleDto)

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
