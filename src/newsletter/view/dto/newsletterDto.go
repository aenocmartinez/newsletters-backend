package dto

type NewsletterDto struct {
	Id             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Section        string `json:"section,omitempty"`
	State          string `json:"state,omitempty"`
	HtmlTemplate   string `json:"html_template,omitempty"`
	Recipients     string `json:"recipients,omitempty"`
	ScheduleStatus string `json:"schedule_status,omitempty"`
	ScheduleDate   string `json:"schedule_date,omitempty"`
	ScheduleHour   string `json:"schedule_hour,omitempty"`
}
