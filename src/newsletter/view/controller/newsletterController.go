package controller

import (
	"net/http"
	"pulzo/src/newsletter/usecase"
	newsdto "pulzo/src/newsletter/view/dto"
	formrequest "pulzo/src/newsletter/view/form-request"
	"pulzo/src/shared/view/dto"
	"strconv"

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
