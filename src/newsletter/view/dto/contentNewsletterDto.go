package dto

type ContentNewsletterDto struct {
	Header   string   `json:"header"`
	Footer   string   `json:"footer"`
	Articles []string `json:"articles"`
}
