package controller

import (
	"net/http"
	"pulzo/src/newsletter/usecase"
	newsdto "pulzo/src/newsletter/view/dto"
	formrequest "pulzo/src/newsletter/view/form-request"
	"pulzo/src/shared/infraestructure/util"
	"pulzo/src/shared/view/dto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateNewsletter(c *gin.Context) {
	var req formrequest.CreateNewsletter
	err := c.ShouldBind(&req)
	if err != nil {
		response := dto.ResponseHttp{
			Code:    "400",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": response})
		return
	}

	useCase := usecase.CreateNewsletterUseCase{}
	response := useCase.Execute(newsdto.NewsletterDto{
		Name:         req.Name,
		Section:      req.Section,
		HtmlTemplate: req.HtmlTemplate,
		Recipients:   req.Recipients,
	})

	code, _ := strconv.Atoi(response.Code)
	if code != 200 && code != 201 {
		c.JSON(code, gin.H{"error": response})
		return
	}

	c.JSON(code, response)
}

func NewsletterList(c *gin.Context) {
	useCase := usecase.NewsletterListUseCase{}
	response := useCase.Execute()
	code, _ := strconv.Atoi(response.Code)

	c.JSON(code, response)
}

func EditNewsletter(c *gin.Context) {
	var strId string = c.Query("id")

	idNewsletter, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		return
	}

	useCase := usecase.EditNewsletterUseCase{}
	response := useCase.Execute(int64(idNewsletter))
	code, _ := strconv.Atoi(response.Code)

	c.JSON(code, response)
}

func UpdateNewsletter(c *gin.Context) {
	var req formrequest.UpdateNewsletter
	err := c.ShouldBind(&req)
	if err != nil {
		response := dto.ResponseHttp{
			Code:    "400",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": response})
		return
	}

	useCase := usecase.UpdateNewsletterUseCase{}
	response := useCase.Execute(newsdto.NewsletterDto{
		Id:           req.Id,
		Name:         req.Name,
		Section:      req.Section,
		HtmlTemplate: req.HtmlTemplate,
		State:        req.State,
		Recipients:   req.Recipients,
	})

	code, _ := strconv.Atoi(response.Code)
	if code != 200 && code != 201 {
		c.JSON(code, gin.H{"error": response})
		return
	}

	c.JSON(code, response)
}

func DeleteNewsletter(c *gin.Context) {
	var req formrequest.DeleteNewsletter
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.DeleteNewsletterUseCase{}
	response := useCase.Execute(req.Id)

	code, _ := strconv.Atoi(response.Code)
	if code != 200 && code != 201 {
		c.JSON(code, gin.H{"error": response})
		return
	}

	c.JSON(code, response)
}

func ScheduleNewsletter(c *gin.Context) {
	var strId string = c.Query("id")

	idNewsletter, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		return
	}

	useCase := usecase.ScheduleNewslletterUseCase{}
	response := useCase.Execute(int64(idNewsletter))
	code, _ := strconv.Atoi(response.Code)

	c.JSON(code, response)
}

func SaveScheduleNewsletter(c *gin.Context) {
	var req formrequest.SaveScheduleNewsletter
	err := c.ShouldBind(&req)
	if err != nil {
		response := dto.ResponseHttp{
			Code:    "400",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": response})
		return
	}

	if !util.DateFormatYYYYMMDD(req.Date) {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ResponseHttp{
			Code:    "400",
			Message: "Formato de fecha no válido, por favor utilizar YYYY-mm-dd",
		}})
		return
	}

	if !util.HourFormatHHMMSS(req.Hour) {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ResponseHttp{
			Code:    "400",
			Message: "Formato de hora no válido, por favor utilizar HH:mm",
		}})
		return
	}

	today := time.Now()
	dateSchedule, _ := time.Parse(util.YYYY_MM_DD, req.Date)
	isEqualDate := (today.Year() == dateSchedule.Year() && today.YearDay() == dateSchedule.YearDay())

	if !isEqualDate && dateSchedule.Before(today) {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ResponseHttp{
			Code:    "400",
			Message: "No se puede programar un newsletter para una fecha anterior a la actual.",
		}})
		return
	}

	hour, _ := time.Parse(util.HH_MM_SS, req.Hour)
	if !util.HourOfPeriodOf15Minutes(hour) {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ResponseHttp{
			Code:    "400",
			Message: "Formato de hora no válido, solo se permiten minutos con periodos de 15, por ejemplo, 0, 15, 30 o 45",
		}})
		return
	}

	if isEqualDate && util.ItIsAnHourBeforeTheCurrent(req.Hour) {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ResponseHttp{
			Code:    "400",
			Message: "Hora no válida, se debe programar el envío de un newsletter para una hora posterior a la actual",
		}})
		return
	}

	useCase := usecase.SaveSettingNewsletterUseCase{}
	response := useCase.Execute(newsdto.ScheduleNewsletterDto{
		Subject: req.Subject,
		Date:    req.Date,
		Hour:    req.Hour,
		Id:      req.Id,
		From:    req.From,
	})

	code, _ := strconv.Atoi(response.Code)
	if code != 200 && code != 201 {
		c.JSON(code, gin.H{"error": response})
		return
	}

	c.JSON(code, response)
}

func AddArticleToNewsletter(c *gin.Context) {
	var q string = c.Query("q")
	var strIdNewsletter string = c.Query("newsletter")
	idNewsletter, ok := util.ValidateIdNumber(strIdNewsletter)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid param"})
		return
	}

	useCase := usecase.AddArticleToNewsletterUseCase{}
	response := useCase.Execute(q, idNewsletter)
	c.JSON(200, response)
}

func RemoveArticleToNewsletter(c *gin.Context) {

	var req formrequest.RemoveArticleToNewsletter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := dto.ResponseHttp{
			Code:    "400",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": response})
		return
	}

	useCase := usecase.RemoveArticleToNewsletterUseCase{}
	response := useCase.Execute(req.IdNewsletter, req.IdArticle)
	c.JSON(200, response)
}
