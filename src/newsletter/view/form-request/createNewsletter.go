package formrequest

type CreateNewsletter struct {
	Name         string `json:"name" binding:"required,max=80"`
	Section      string `json:"section" binding:"required,max=80"`
	HtmlTemplate string `json:"html_template" binding:"required,max=80"`
	Recipients   string `json:"recipients" binding:"required,max=80"`
}
