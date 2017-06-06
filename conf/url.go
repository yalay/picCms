package conf

import "strconv"

func GetArticleUrl(articleId int32) string {
	return "/article-" + strconv.Itoa(int(articleId)) + ".html"
}

func GetArticlePageUrl(articleArgs ...string) string {
	pageUrl := "/article"
	for _, articleArg := range articleArgs {
		pageUrl += "-" + articleArg
	}
	return pageUrl + ".html"
}

func GetCateUrl(engName string) string {
	return "/" + engName + ".html"
}

func GetTopicUrl(engName string) string {
	return "/topic-" + engName + ".html"
}

func GetTagUrl(tag string) string {
	return "/tags-" + tag + ".html"
}
