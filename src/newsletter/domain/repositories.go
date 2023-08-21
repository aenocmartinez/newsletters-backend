package domain

import "pulzo/src/newsletter/view/dto"

type NewsletterRepository interface {
	CreateNewsletter(newsletter Newsletter) error
	UpdateNewsletter(newsletter Newsletter) error
	DeleteNewsletter(newsletter Newsletter) error
	FindNewsletterById(id int64) Newsletter
	FindNewsletterByName(name string) Newsletter
	NewsletterList() []dto.NewsletterDto
	DeleteSchedule(idNewsletter int64) error
	SaveSchedule(newsletter Newsletter) error
}
