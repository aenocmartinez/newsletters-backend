package usecase

import (
	"pulzo/src/newsletter/dao/mysql"
	"pulzo/src/newsletter/domain"
)

var newsletterRepository domain.NewsletterRepository = mysql.NewNewsletterDao()
