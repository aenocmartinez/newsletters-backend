package domain

import (
	"os"
	"pulzo/src/newsletter/view/dto"
	"strings"
)

type Newsletter struct {
	id                   int64
	name                 string
	section              string
	state                string
	htmlTemplate         string
	recipients           string
	schedule             Schedule
	newsletterRepository NewsletterRepository
}

func NewNewsletter(name string) *Newsletter {
	newsletter := &Newsletter{}
	newsletter.SetName(name)
	newsletter.SetState("active")
	newsletter.SetHtmlTemplate("")

	return newsletter
}

func (n *Newsletter) SetRepository(newsletterRepository NewsletterRepository) {
	n.newsletterRepository = newsletterRepository
}

func (n *Newsletter) SetSchedule(schedule Schedule) {
	n.schedule = schedule
}

func (n *Newsletter) Schedule() Schedule {
	return n.schedule
}

func (n *Newsletter) SetCampaignId(idCampaign string) {
	n.schedule.SetCampaignId(idCampaign)
}

func (n *Newsletter) CampaignId() string {
	return n.schedule.idCampaign
}

func (n *Newsletter) ScheduleDate() string {
	return n.schedule.Date()
}

func (n *Newsletter) ScheduleHour() string {
	return n.schedule.Hour()
}

func (n *Newsletter) ScheduleMessage() string {
	return n.schedule.Message()
}

func (n *Newsletter) ScheduleSubject() string {
	subject := strings.ReplaceAll(n.schedule.Subject(), "\"", "'")
	return subject
}

func (n *Newsletter) ScheduleFrom() string {
	return n.schedule.From()
}

func (n *Newsletter) ScheduleStatus() string {
	return n.schedule.Status()
}

func (n *Newsletter) ScheduleWasExecuted() bool {
	return n.schedule.WasExecuted()
}

func (n *Newsletter) SetId(id int64) {
	n.id = id
}

func (n *Newsletter) Id() int64 {
	return n.id
}

func (n *Newsletter) SetSection(section string) {
	if len(section) > 80 {
		section = section[:80]
	}
	n.section = section
}

func (n *Newsletter) Section() string {
	return n.section
}

func (n *Newsletter) SetName(name string) {
	if len(name) > 80 {
		name = name[:80]
	}
	n.name = name
}

func (n *Newsletter) Name() string {
	return n.name
}

func (n *Newsletter) SetState(state string) {
	if len(state) == 0 {
		state = "active"
	}

	if !strings.Contains("active,inactive", strings.ToLower(state)) {
		state = "inactive"
	}
	n.state = state
}

func (n *Newsletter) State() string {
	return strings.ToLower(n.state)
}

func (n *Newsletter) SetHtmlTemplate(htmlTemplate string) {
	if len(htmlTemplate) == 0 {
		htmlTemplate = os.Getenv("HTML_TEMPLATE_DEFAULT")
	}
	n.htmlTemplate = htmlTemplate
}

func (n *Newsletter) HtmlTemplate() string {
	return n.htmlTemplate
}

func (n *Newsletter) SetRecipients(recipients string) {
	n.recipients = recipients
}

func (n *Newsletter) Recipients() string {
	return n.recipients
}

func (n *Newsletter) Create() error {
	return n.newsletterRepository.CreateNewsletter(*n)
}

func (n *Newsletter) Update() error {
	return n.newsletterRepository.UpdateNewsletter(*n)
}

func (n *Newsletter) Delete() error {
	return n.newsletterRepository.DeleteNewsletter(*n)
}

func NewsletterList(newsletterRepository NewsletterRepository) []dto.NewsletterDto {
	return newsletterRepository.NewsletterList()
}

func FindNewsletterById(id int64, newsletterRepository NewsletterRepository) Newsletter {
	return newsletterRepository.FindNewsletterById(id)
}

func FindNewsletterByName(name string, newsletterRepository NewsletterRepository) Newsletter {
	return newsletterRepository.FindNewsletterByName(name)
}

func (n *Newsletter) Exists() bool {
	return n.id > 0
}

func (n *Newsletter) IsMailchimp() bool {
	return strings.ToLower(os.Getenv("APP_SENT_BY")) == "mailchimp"
}

func (n *Newsletter) ScheduleDelivery() error {
	err := n.newsletterRepository.DeleteSchedule(n.id)
	if err != nil {
		return err
	}
	return n.newsletterRepository.SaveSchedule(*n)
}

func (n *Newsletter) DeleteSchedule() {
	n.newsletterRepository.DeleteSchedule(n.id)
}

func (n *Newsletter) HasSchedule() bool {
	return n.schedule.Exists()
}

func (n *Newsletter) JsonFileName() string {
	return strings.ReplaceAll(strings.ToLower(n.name), " ", "_")
}
