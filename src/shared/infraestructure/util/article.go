package util

import "strings"

func ExtractIdArticleFromURL(url string) string {
	result := strings.Split(url, "-PP")
	if len(result) > 1 {
		return "PP" + result[1]
	}
	return ""
}
