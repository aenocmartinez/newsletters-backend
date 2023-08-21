package domain

import (
	"encoding/json"
	"log"
	"os"
	"pulzo/src/newsletter/view/dto"
	"pulzo/src/shared/infraestructure/util"
	"strconv"
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
	arrayArticle         []dto.ArticleDto
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

func (n *Newsletter) AddArticle(articleDto dto.ArticleDto) {
	var newArrayArticle []dto.ArticleDto = []dto.ArticleDto{}
	n.ReadFileJson()

	for _, item := range n.arrayArticle {
		if item.Link == articleDto.Link {
			continue
		}
		newArrayArticle = append(newArrayArticle, item)
	}

	newArrayArticle = append([]dto.ArticleDto{articleDto}, newArrayArticle...)
	n.arrayArticle = newArrayArticle
}

func (n *Newsletter) RemoveArticle(posItem int) {
	var newArrayArticle []dto.ArticleDto = []dto.ArticleDto{}
	n.ReadFileJson()

	for index, item := range n.arrayArticle {
		if index == posItem {
			continue
		}
		newArrayArticle = append(newArrayArticle, item)
	}
	n.arrayArticle = newArrayArticle
}

func (n *Newsletter) ReadFileJson() {
	var arrayArticle []dto.ArticleDto = []dto.ArticleDto{}
	if n.schedule.WasExecuted() {
		n.arrayArticle = arrayArticle
		return
	}

	urlJson := os.Getenv("TEMPLATE_PATH") + "/" + n.section + "/" + n.JsonFileName() + ".json"
	content, err := util.ReadFileFromURL(urlJson)
	if err == nil {
		json.Unmarshal(content, &arrayArticle)
	}

	n.arrayArticle = arrayArticle
}

func (n *Newsletter) ArrayArticle() (arrayArticle []dto.ArticleDto) {
	return n.arrayArticle
}

func (n *Newsletter) Setting() (content dto.ContentNewsletterDto) {
	footer := n.getFooter()
	header := n.getHeader()

	contentHeader := "<table align=\"center\" style=\"border-spacing: 0; font-family: 'Heebo', sans-serif; color: #333333; margin: 0 auto; width: 100%; max-width: 700px;\">"
	contentHeader += "<tbody>"
	contentHeader += header

	for index, article := range n.arrayArticle {
		if index == 0 {
			var fItem string = n.getFirstArticle(article)
			fItem = strings.ReplaceAll(fItem, "\n", "")
			fItem = strings.ReplaceAll(fItem, "\\", "")
			content.Articles = append(content.Articles, fItem)
			continue
		}

		var item string = n.getOtherArticle(article, index+1)
		item = strings.ReplaceAll(item, "\n", "")
		item = strings.ReplaceAll(item, "\\", "")
		content.Articles = append(content.Articles, item)
	}

	contentFooter := footer
	contentFooter += "</tbody>"
	contentFooter += "</table>"

	contentHeader = strings.ReplaceAll(contentHeader, "\n", "")
	contentHeader = strings.ReplaceAll(contentHeader, "\\", "")

	contentFooter = strings.ReplaceAll(contentFooter, "\n", "")
	contentFooter = strings.ReplaceAll(contentFooter, "\\", "")

	content.Header = contentHeader
	content.Footer = contentFooter

	return content
}

func (n *Newsletter) GenerateHtml() (html string) {
	header := n.getHeader()
	footer := n.getFooter()

	html += "<tbody>"
	html += header

	for index, article := range n.arrayArticle {
		if index == 0 {
			html += n.getFirstArticle(article)
			continue
		}
		html += n.getOtherArticle(article, index+1)
	}

	html += footer

	html += "</tbody>"
	html += "</table>"

	html = strings.ReplaceAll(html, "'", "")
	html = strings.ReplaceAll(html, "\"", "'")
	html = strings.ReplaceAll(html, "\n", "")
	html = strings.ReplaceAll(html, "\t", "")
	html = strings.ReplaceAll(html, "\\", "")

	return html
}

func (n *Newsletter) getHeader() string {
	urlHeader := os.Getenv("TEMPLATE_PATH") + "/templates_html/" + n.htmlTemplate + "/header.html"
	contentHeader, err := util.ReadFileFromURL(urlHeader)
	if err != nil {
		log.Println(err)
	}
	header := string(contentHeader)

	header = strings.ReplaceAll(header, "{{SECTION}}", strings.ToUpper(n.section))
	header = strings.ReplaceAll(header, "{{DATE}}", util.DateFormatForNewsletter())

	return header
}

func (n *Newsletter) getFooter() string {
	urlFooter := os.Getenv("TEMPLATE_PATH") + "/templates_html/" + n.htmlTemplate + "/footer.html"
	contentFooter, err := util.ReadFileFromURL(urlFooter)
	if err != nil {
		log.Println(err)
	}

	footer := string(contentFooter)
	footer = strings.ReplaceAll(footer, "{{SECTION}}", strings.ToUpper(n.section))
	footer = strings.ReplaceAll(footer, "{{SECTION_LINK}}", n.section)

	return string(footer)
}

func (n *Newsletter) getFirstArticle(articleDto dto.ArticleDto) string {

	urlFirstArticle := os.Getenv("TEMPLATE_PATH") + "/templates_html/" + n.htmlTemplate + "/first_article.html"
	contentFirstArticle, err := util.ReadFileFromURL(urlFirstArticle)
	if err != nil {
		log.Println(err)
	}

	firstArticle := string(contentFirstArticle)
	firstArticle = strings.ReplaceAll(firstArticle, "{{SECTION}}", strings.ToUpper(n.section))
	firstArticle = strings.ReplaceAll(firstArticle, "{{TITLE}}", articleDto.Title)
	firstArticle = strings.ReplaceAll(firstArticle, "{{LEAD}}", articleDto.Lead)
	firstArticle = strings.ReplaceAll(firstArticle, "{{LINK}}", "https://www.pulzo.com"+articleDto.Link)
	firstArticle = strings.ReplaceAll(firstArticle, "{{IMAGE}}", articleDto.ImagenFA)
	firstArticle = strings.ReplaceAll(firstArticle, "\n", "")
	firstArticle = strings.ReplaceAll(firstArticle, "\\", "")

	return firstArticle
}

func (n *Newsletter) getOtherArticle(articleDto dto.ArticleDto, index int) string {
	urlItem := os.Getenv("TEMPLATE_PATH") + "/templates_html/" + n.htmlTemplate + "/item.html"

	contentArticle, err := util.ReadFileFromURL(urlItem)
	if err != nil {
		return ""
	}

	article := string(contentArticle)
	article = strings.ReplaceAll(article, "{{TITLE}}", articleDto.Title)
	article = strings.ReplaceAll(article, "{{LINK}}", "https://www.pulzo.com"+articleDto.Link)
	article = strings.ReplaceAll(article, "{{IMAGE}}", articleDto.Image)
	article = strings.ReplaceAll(article, "{{LEAD}}", articleDto.Lead)
	article = strings.ReplaceAll(article, "\n", "")
	article = strings.ReplaceAll(article, "\\", "")
	if n.htmlTemplate == "noticia" {
		article = strings.ReplaceAll(article, "{{INDEX}}", strconv.Itoa(index))
	}

	return article
}
