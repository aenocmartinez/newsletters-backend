package mock

import (
	"errors"
	"pulzo/src/newsletter/domain"
	"pulzo/src/newsletter/view/dto"
	"strconv"
	"strings"
)

type NewsletterDao struct {
	arrayNewsletter []domain.Newsletter
}

func NewNewsletterDao() *NewsletterDao {
	return &NewsletterDao{
		arrayNewsletter: initializeNewsletterData(),
	}
}

func initializeNewsletterData() (arrayNewsletter []domain.Newsletter) {
	for i := 1; i <= 10; i++ {

		strId := strconv.Itoa(i)

		newsletter := domain.NewNewsletter("")
		newsletter.SetId(int64(i))
		newsletter.SetName("Name newsletter " + strId)
		newsletter.SetSection("Section_" + strId)
		newsletter.SetState("active")

		arrayNewsletter = append(arrayNewsletter, *newsletter)
	}

	return arrayNewsletter
}

func (n *NewsletterDao) CreateNewsletter(newsletter domain.Newsletter) error {

	isUpdate := false
	err := validateInput(newsletter, isUpdate)
	if err != nil {
		return err
	}

	newId := int64(len(n.arrayNewsletter) + 1)
	newsletter.SetId(newId)

	n.arrayNewsletter = append(n.arrayNewsletter, newsletter)
	return nil
}

func (n *NewsletterDao) UpdateNewsletter(newsletter domain.Newsletter) error {
	isUpdate := true
	err := validateInput(newsletter, isUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (n *NewsletterDao) DeleteNewsletter(newsletter domain.Newsletter) error {
	if newsletter.Id() == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (n *NewsletterDao) FindNewsletterById(id int64) domain.Newsletter {
	for _, item := range n.arrayNewsletter {
		if item.Id() == id {
			return item
		}
	}
	return domain.Newsletter{}
}

func (n *NewsletterDao) FindNewsletterByName(name string) domain.Newsletter {
	name = strings.ToLower(name)
	for _, item := range n.arrayNewsletter {
		if strings.ToLower(item.Name()) == name {
			return item
		}
	}
	return domain.Newsletter{}
}

func (n *NewsletterDao) NewsletterList() (arrayNewsletter []dto.NewsletterDto) {
	for _, item := range n.arrayNewsletter {
		arrayNewsletter = append(arrayNewsletter, dto.NewsletterDto{
			Id:    item.Id(),
			Name:  item.Name(),
			State: item.State(),
		})
	}
	return arrayNewsletter
}

func (n *NewsletterDao) AddArticleToNewsletter(newsletter domain.Newsletter, article dto.ArticleDto) error {
	if newsletter.Id() == 0 {
		return errors.New("idNewsletter is Zero(0)")
	}
	if len(article.Id) == 0 {
		return errors.New("idArticle is Zero(0)")
	}
	return nil
}

func (n *NewsletterDao) RemoveArticleToNewsletter(newsletter domain.Newsletter, article dto.ArticleDto) error {
	if newsletter.Id() == 0 {
		return errors.New("idNewsletter is Zero(0)")
	}
	if len(article.Id) == 0 {
		return errors.New("idArticle is Zero(0)")
	}
	return nil
}

func (n *NewsletterDao) DeleteSchedule(idNewsletter int64) error {
	return nil
}

func (n *NewsletterDao) SaveSchedule(newsletter domain.Newsletter) error {
	return nil
}

func validateInput(newsletter domain.Newsletter, isUpdate bool) error {

	if isUpdate {
		if newsletter.Id() == 0 {
			return errors.New("id is not valid")
		}
	}

	if len(newsletter.Name()) == 0 {
		return errors.New("name is NULL")
	}

	if len(newsletter.Section()) == 0 {
		return errors.New("section is NULL")
	}

	if len(newsletter.State()) == 0 {
		return errors.New("state is NULL")
	}

	if len(newsletter.HtmlTemplate()) == 0 {
		return errors.New("htmlTemplate is NULL")
	}

	if !strings.Contains("active,inactive", newsletter.State()) {
		return errors.New("state is not valid")
	}

	return nil
}
