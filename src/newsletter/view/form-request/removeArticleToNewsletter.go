package formrequest

type RemoveArticleToNewsletter struct {
	IdNewsletter int64 `json:"idNewsletter" binding:"required"`
	IdArticle    int   `json:"idArticle"`
}
