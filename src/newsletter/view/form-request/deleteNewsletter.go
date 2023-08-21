package formrequest

type DeleteNewsletter struct {
	Id int64 `json:"id" binding:"required"`
}
