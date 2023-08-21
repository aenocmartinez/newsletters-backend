package domain

type ArticleRepository interface {
	FindArticleById(id string) Article
}
