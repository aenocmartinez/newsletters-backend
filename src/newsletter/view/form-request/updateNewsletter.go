package formrequest

type UpdateNewsletter struct {
	Id           int64  `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required,max=80"`
	Section      string `json:"section" binding:"required,max=80"`
	HtmlTemplate string `json:"html_template" binding:"required,max=250"`
	State        string `json:"state"`
	Recipients   string `json:"recipients" binding:"required,max=80"`
}
