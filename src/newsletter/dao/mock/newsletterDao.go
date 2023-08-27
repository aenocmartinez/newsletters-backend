package mock

import (
	"errors"
	"pulzo/src/newsletter/domain"
	"pulzo/src/newsletter/view/dto"
	"strings"
)

type NewsletterDao struct {
	data []domain.Newsletter
}

func NewNewsletterDao() *NewsletterDao {
	dao := &NewsletterDao{}
	dao.loadData()

	return dao
}

func (n *NewsletterDao) loadData() {
	var dataDefault []string = []string{"entretenimiento", "noticias", "deportes"}

	for index, item := range dataDefault {
		newsletter := domain.NewNewsletter(item)
		newsletter.SetSection(strings.ToUpper(item))
		newsletter.SetRecipients("11605")
		newsletter.SetId(int64(index + 1))
		n.data = append(n.data, *newsletter)
	}

}

func (n *NewsletterDao) CreateNewsletter(newsletter domain.Newsletter) error {
	newsletterFound := n.FindNewsletterByName(newsletter.Name())
	if newsletterFound.Exists() {
		return errors.New("error")
	}

	nextId := len(n.data) + 1
	newsletter.SetId(int64(nextId))
	n.data = append(n.data, newsletter)

	return nil
}

func (n *NewsletterDao) UpdateNewsletter(newsletter domain.Newsletter) error {
	data := []domain.Newsletter{}
	var flag bool = false

	newsletterFound := n.FindNewsletterByName(newsletter.Name())
	if newsletterFound.Exists() {
		return errors.New("error")
	}

	for _, item := range n.data {
		if item.Id() == newsletter.Id() {
			item.SetName(newsletter.Name())
			item.SetSection(newsletter.Section())
			item.SetState(newsletter.State())
			item.SetHtmlTemplate(newsletter.HtmlTemplate())
			item.SetRecipients(newsletter.Recipients())
			data = append(data, item)
			flag = true
		}
	}

	if !flag {
		return errors.New("error")
	}

	n.data = data
	return nil
}

func (n *NewsletterDao) DeleteNewsletter(newsletter domain.Newsletter) error {
	data := []domain.Newsletter{}
	var flag bool = false
	for _, item := range n.data {
		if item.Id() == newsletter.Id() {
			flag = true
			continue
		}
		data = append(data, item)
	}

	if !flag {
		return errors.New("error")
	}

	n.data = data

	return nil
}

func (n *NewsletterDao) FindNewsletterById(id int64) domain.Newsletter {
	for _, item := range n.data {
		if item.Id() == id {
			return item
		}
	}
	return domain.Newsletter{}
}

func (n *NewsletterDao) FindNewsletterByName(name string) domain.Newsletter {
	for _, item := range n.data {
		if strings.EqualFold(item.Name(), name) {
			return item
		}
	}
	return domain.Newsletter{}
}

func (n *NewsletterDao) NewsletterList() []dto.NewsletterDto {
	list := []dto.NewsletterDto{}
	for _, item := range n.data {
		list = append(list, dto.NewsletterDto{
			Id:             item.Id(),
			Name:           item.Name(),
			Section:        item.Section(),
			State:          item.State(),
			HtmlTemplate:   item.HtmlTemplate(),
			Recipients:     item.Recipients(),
			ScheduleStatus: item.ScheduleStatus(),
			ScheduleDate:   item.ScheduleDate(),
			ScheduleHour:   item.ScheduleHour(),
		})
	}
	return list
}

func (n *NewsletterDao) DeleteSchedule(idNewsletter int64) error {
	data := []domain.Newsletter{}
	var flag bool = false
	for _, item := range n.data {
		if item.Id() == idNewsletter {
			item.SetSchedule(domain.Schedule{})
			flag = true
			continue
		}
		data = append(data, item)
	}

	if !flag {
		return errors.New("error")
	}

	n.data = data

	return nil
}

func (n *NewsletterDao) SaveSchedule(newsletter domain.Newsletter) error {
	data := []domain.Newsletter{}
	var flag bool = false
	for _, item := range n.data {
		if item.Id() == newsletter.Id() {
			item.SetSchedule(newsletter.Schedule())
			data = append(data, item)
			flag = true
		}
	}

	if !flag {
		return errors.New("error")
	}

	n.data = data
	return nil
}
