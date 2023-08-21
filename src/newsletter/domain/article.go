package domain

import (
	"pulzo/src/newsletter/view/dto"
	"strings"
)

func FindArticleById(id string, articleRepository ArticleRepository) dto.ArticleDto {
	return articleRepository.FindArticleById(id)
}

func ExtractIdArticleFromURL(url string) string {
	result := strings.Split(url, "-PP")
	if len(result) > 1 {
		return "PP" + result[1]
	}
	return ""
}
