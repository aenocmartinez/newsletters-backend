package mysql

import (
	"bytes"
	"errors"
	"pulzo/src/newsletter/domain"
	"pulzo/src/newsletter/view/dto"
	"pulzo/src/shared/infraestructure/database"
	"strings"

	"github.com/getsentry/sentry-go"
)

type NewsletterDao struct{}

func NewNewsletterDao() *NewsletterDao {
	return &NewsletterDao{}
}

func (n *NewsletterDao) CreateNewsletter(newsletter domain.Newsletter) error {
	isUpdate := false
	err := validateInput(newsletter, isUpdate)
	if err != nil {
		return err
	}

	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("INSERT INTO newsletters(name, section, state, htmlTemplate, recipients) VALUES (?, ?, ?, ?, ?)")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	_, err = stmt.Exec(newsletter.Name(), newsletter.Section(), newsletter.State(), newsletter.HtmlTemplate(), newsletter.Recipients())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (n *NewsletterDao) UpdateNewsletter(newsletter domain.Newsletter) error {
	isUpdate := true
	err := validateInput(newsletter, isUpdate)
	if err != nil {
		return err
	}

	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("UPDATE newsletters SET name=?, section=?, state=?, htmlTemplate=?, updatedAt=NOW(), recipients=? WHERE id=?")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	_, err = stmt.Exec(newsletter.Name(), newsletter.Section(), newsletter.State(), newsletter.HtmlTemplate(), newsletter.Recipients(), newsletter.Id())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (n *NewsletterDao) DeleteNewsletter(newsletter domain.Newsletter) error {
	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("DELETE FROM newsletters WHERE id=?")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	_, err = stmt.Exec(newsletter.Id())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (n *NewsletterDao) FindNewsletterById(idNewsletter int64) (newsletter domain.Newsletter) {
	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, name, section, state, htmlTemplate, recipients FROM newsletters WHERE id=?")

	rs := db.Conn().QueryRow(strSQL.String(), idNewsletter)

	var id int64
	var name, section, state, htmlTemplate, recipients string

	rs.Scan(&id, &name, &section, &state, &htmlTemplate, &recipients)

	newsletter = *domain.NewNewsletter(name)
	newsletter.SetId(id)
	newsletter.SetSection(section)
	newsletter.SetState(state)
	newsletter.SetHtmlTemplate(htmlTemplate)
	newsletter.SetRecipients(recipients)
	newsletter.SetSchedule(findSchedule(id))

	return newsletter
}

func (n *NewsletterDao) FindNewsletterByName(name string) (newsletter domain.Newsletter) {
	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, name, section, state, htmlTemplate, recipients FROM newsletters WHERE name = ?")

	rs := db.Conn().QueryRow(strSQL.String(), name)

	var id int64
	var section, state, htmlTemplate, recipients string

	rs.Scan(&id, &name, &section, &state, &htmlTemplate, &recipients)

	newsletter = *domain.NewNewsletter(name)
	newsletter.SetId(id)
	newsletter.SetSection(section)
	newsletter.SetState(state)
	newsletter.SetHtmlTemplate(htmlTemplate)
	newsletter.SetRecipients(recipients)

	return newsletter
}

func (n *NewsletterDao) NewsletterList() (arrayNewsletter []dto.NewsletterDto) {

	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT id, name, section, state, htmlTemplate, recipients FROM newsletters ORDER BY name")

	db := database.InstanceDB()

	rows, err := db.Conn().Query(strQuery.String())
	if err != nil {
		sentry.CaptureException(err)
		return arrayNewsletter
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, section, state, htmlTemplate, recipients string

		rows.Scan(&id, &name, &section, &state, &htmlTemplate, &recipients)

		schedule := findSchedule(id)

		arrayNewsletter = append(arrayNewsletter, dto.NewsletterDto{
			Id:             id,
			Name:           name,
			Section:        section,
			State:          state,
			HtmlTemplate:   htmlTemplate,
			Recipients:     recipients,
			ScheduleStatus: schedule.Status(),
			ScheduleDate:   schedule.Date(),
			ScheduleHour:   schedule.Hour(),
		})
	}

	return arrayNewsletter
}

func (n *NewsletterDao) DeleteSchedule(idNewsletter int64) error {
	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("DELETE FROM schedule WHERE newsletter_id = ?")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	_, err = stmt.Exec(idNewsletter)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func (n *NewsletterDao) SaveSchedule(newsletter domain.Newsletter) error {
	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("INSERT INTO schedule (newsletter_id, campaign_id, subject, send_from, send_date, hour, message) VALUES (?, ?, ?, ?, ?, ?, ?)")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	_, err = stmt.Exec(newsletter.Id(), newsletter.CampaignId(), newsletter.ScheduleSubject(), newsletter.ScheduleFrom(), newsletter.ScheduleDate(), newsletter.ScheduleHour(), newsletter.ScheduleMessage())
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func findSchedule(idNewsletter int64) (schedule domain.Schedule) {
	db := database.InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT campaign_id, send_date, hour, message, subject FROM schedule WHERE newsletter_id = ?")

	rs := db.Conn().QueryRow(strSQL.String(), idNewsletter)

	var campaign, date, hour, message, subject string

	rs.Scan(&campaign, &date, &hour, &message, &subject)

	if len(hour) > 0 {
		hour = hour[:5]
	}

	schedule = domain.Schedule{}
	schedule.SetCampaignId(campaign)
	schedule.SetDate(date)
	schedule.SetHour(hour)
	schedule.SetMessage(message)
	schedule.SetSubject(subject)

	return schedule
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
