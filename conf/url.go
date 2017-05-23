package conf

import "strconv"

func GetArticleUrl(articleId int32) string {
	return "/article-" + strconv.Itoa(int(articleId)) + ".html"
}

func GetCateUrl(engName string) string {
	return "/" + engName + ".html"
}

func GetTagUrl(tag string) string  {
	return "/tags-" + tag + ".html"
}
