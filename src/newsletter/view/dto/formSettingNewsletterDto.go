package dto

type FormSettingNewsletterDto struct {
	IdNewsletter int64                `json:"id"`
	Date         string               `json:"date,omitempty"`
	Hour         string               `json:"hour,omitempty"`
	Html         string               `json:"html,omitempty"`
	Content      ContentNewsletterDto `json:"content"`
	Status       string               `json:"status,omitempty"`
	Subject      string               `json:"subject,omitempty"`
}
