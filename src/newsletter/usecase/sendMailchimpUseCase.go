package usecase

import (
	"pulzo/src/newsletter/domain"
	pulzoservices "pulzo/src/newsletter/infraestructure/pulzo_services"

	"github.com/getsentry/sentry-go"
)

type SendMailchimpUseCase struct{}

func (useCase *SendMailchimpUseCase) Execute(newsletter domain.Newsletter) {

	messenger := pulzoservices.NewMessenger(&newsletter)
	message := messenger.Execute()

	if len(message) > 0 {
		sentry.CaptureMessage(message)
	}

	err := newsletter.ScheduleDelivery()
	if err != nil {
		sentry.CaptureException(err)
	}
}
