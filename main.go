package main

import (
	"log"
	"os"

	"pulzo/src/newsletter/view/controller"

	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY"),
	})

	if err != nil {
		log.Println("Error al iniciar sentry: ", err)
	}

	defer sentry.Flush(2 * time.Second)

	func() {
		defer func() {
			err := recover()
			if err != nil {
				sentry.CurrentHub().Recover(err)
				sentry.Flush(time.Second * 5)
			}
		}()

		r := gin.New()
		r.Use(gin.Logger(), gin.Recovery())

		r.POST("newsletters", controller.CreateNewsletter)
		r.GET("newsletters", controller.NewsletterList)
		r.GET("newsletters-edit", controller.EditNewsletter)
		r.PUT("newsletters", controller.UpdateNewsletter)
		r.DELETE("newsletters", controller.DeleteNewsletter)

		r.Run(":8085")

	}()
}
