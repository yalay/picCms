package conf

import "strconv"

func GetArticleUrl(articleId int32) string {
	return "/post-" + strconv.Itoa(int(articleId)) + ".html"
}

func GetCateUrl(cateId int32) string {
	return "/list-" + strconv.Itoa(int(cateId)) + ".html"
}
