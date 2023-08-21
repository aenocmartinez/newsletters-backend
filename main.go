package main

import (
	"log"
	"os"

	"pulzo/src/newsletter/view/controller"

	"pulzo/src/shared/infraestructure/database"
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

		// gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(gin.Logger(), gin.Recovery())

		r.GET("/mutant", func(c *gin.Context) {
			c.String(200, "Redaccion up\n")
		})

		r.GET("/mutant-db", func(c *gin.Context) {
			db := database.InstanceDB()
			err := db.Conn().Ping()
			if err != nil {
				c.String(200, "Error conexion db 2\n")
			}

			c.String(200, "Inventario-db up\n")
		})

		// Newsletters
		r.POST("newsletters", controller.CreateNewsletter)
		r.GET("newsletters", controller.NewsletterList)
		r.GET("newsletters-edit", controller.EditNewsletter)
		r.PUT("newsletters", controller.UpdateNewsletter)
		r.DELETE("newsletters", controller.DeleteNewsletter)
		r.GET("newsletters-schedule", controller.ScheduleNewsletter)
		r.POST("newsletters-schedule", controller.SaveScheduleNewsletter)
		r.GET("newsletters-add-article", controller.AddArticleToNewsletter)
		r.DELETE("newsletters-remove-article", controller.RemoveArticleToNewsletter)

		r.Run(":8085")

	}()
}
