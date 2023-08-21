package controller_test

import (
	"net/http"
	"net/http/httptest"
	"pulzo/src/newsletter/usecase"
	formrequest "pulzo/src/newsletter/view/form-request"
	"pulzo/src/shared/infraestructure/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

func TestValidateInputCreateNewsletter(t *testing.T) {
	c := GetTestGinContext()
	req := formrequest.CreateNewsletter{
		Name: "Lorem ipsum dolor sit amet consectetur adipiscing elit dapibus mus suspendisse quisque suscipit",
	}
	err := c.ShouldBind(&req)
	if err != nil {
		t.Log("ok")
	} else {
		t.Error(err)
	}
}

func TestIdParamDeleteNewsletter(t *testing.T) {
	c := GetTestGinContext()
	c.AddParam("id", "30")
	var id string = c.Param("id")
	var expectedValue bool = true
	_, result := util.ValidateIdNumber(id)
	assert.Equal(t, result, expectedValue)
}

func TestDeleteNewsletter(t *testing.T) {
	c := GetTestGinContext()
	c.AddParam("id", "3000")
	var strId string = c.Param("id")

	id, result := util.ValidateIdNumber(strId)
	if !result {
		t.Fatal("unexpected error")
	}

	useCase := usecase.DeleteNewsletterUseCase{}
	response := useCase.Execute(id)
	if response.Code != "404" {
		t.Fatal("unexpected error")
	}
}
