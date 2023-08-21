package pulzoservices

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"pulzo/src/newsletter/domain"
	"pulzo/src/shared/infraestructure/util"

	"strings"

	"github.com/getsentry/sentry-go"
)

type Messenger struct {
	newsletter *domain.Newsletter
}

type template struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Id    string `json:"id"`
}

func (t *template) getId(newsletter domain.Newsletter) string {
	arrayTemplate := []template{}
	urlTemplateJson := "https://filesstaticpulzo.s3.us-west-2.amazonaws.com/pulzo-lite/newsletters/templates.json"
	content, err := util.ReadFileFromURL(urlTemplateJson)
	if err != nil {
		return "10000008"
	}

	json.Unmarshal(content, &arrayTemplate)

	for _, template := range arrayTemplate {
		if template.Value == newsletter.HtmlTemplate() {
			return template.Id
		}
	}
	return "10000008"
}

func NewMessenger(newsletter *domain.Newsletter) *Messenger {
	return &Messenger{
		newsletter: newsletter,
	}
}

func (m *Messenger) Execute() string {
	url := os.Getenv("ENDPOINT_PULZO_MESSENGER")
	sent_by := os.Getenv("APP_SENT_BY")

	recipients := m.newsletter.Recipients()
	if strings.ToLower(sent_by) == "mailtrap" {
		recipients = os.Getenv("APP_TEST_RECIPIENTS")
	}

	recipients = strings.ReplaceAll(recipients, " ", "")
	arrayRecipients := strings.Split(recipients, ",")
	recipients = ""
	for _, recipient := range arrayRecipients {
		recipients += "\"" + recipient + "\","
	}

	recipients = recipients[:len(recipients)-1]

	template := template{}
	idTemplate := template.getId(*m.newsletter)

	method := "POST"

	payload := strings.NewReader(`{
	"message":"` + m.newsletter.ScheduleMessage() + `",
	"subject":"` + m.newsletter.ScheduleSubject() + `",
	"from":"` + m.newsletter.ScheduleFrom() + `",
	"recipients":[` + recipients + `],
	"date":"` + m.newsletter.ScheduleDate() + `",
	"hour":"` + m.newsletter.ScheduleHour() + `",
	"idCampaign":"` + m.newsletter.CampaignId() + `",
	"idTemplate":"` + idTemplate + `",
	"sent_by":"` + sent_by + `"
  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		sentry.CaptureException(err)
		return err.Error()
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		sentry.CaptureException(err)
		return err.Error()
	}

	response := string(body)

	if !strings.Contains(response, "success") {
		return response
	}

	if strings.ToLower(sent_by) == "mailchimp" {
		pos := strings.Index(response, "idCampaign")
		idCampaign := response[pos+len("idCampaign")+3:]
		idCampaign = strings.ReplaceAll(idCampaign, "}", "")
		idCampaign = strings.ReplaceAll(idCampaign, "\"", "")
		m.newsletter.SetCampaignId(idCampaign)
	}

	return ""
}
