package dto

type ScheduleNewsletterDto struct {
	Id      int64
	Subject string
	Date    string
	Hour    string
	From    string
}

type ArticleDto struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Lead     string `json:"lead"`
	Link     string `json:"link"`
	Section  string `json:"section"`
	Image    string `json:"image"`
	ImagenFA string `json:"imagenFA"`
}
