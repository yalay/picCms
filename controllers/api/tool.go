package api

import (
	"controllers"
	"strings"
	"sync"
	"time"

	"fmt"
	"github.com/astaxie/beego/context"
)

func ToolHandler(c *context.Context) {
	toolType := c.Input.Param(":type")
	id := c.Input.Param(":id")
	switch toolType {
	case "translate":
		articleId := Atoi32(id)
		if articleId == 0 {
			c.WriteString("no article id")
			return
		}
		article := controllers.GetArticle(articleId)
		if article == nil {
			c.WriteString("invalid article id")
			return
		}

		// 采用#分割
		oriTexts := article.Title + "#" + strings.Replace(article.Keywords, ",", "#", -1)
		var wg sync.WaitGroup
		var chtTexts, engTexts string
		go func() {
			wg.Add(1)
			chtTexts, _ = controllers.TranslateToCht(oriTexts)
			wg.Done()
		}()
		go func() {
			wg.Add(1)
			engTexts, _ = controllers.TranslateToEng(oriTexts)
			wg.Done()
		}()
		time.Sleep(10 * time.Millisecond)
		wg.Wait()

		oriTextFields := strings.Split(oriTexts, "#")
		chtTextFields := strings.Split(chtTexts, "#")
		engTextFields := strings.Split(engTexts, "#")

		var result string
		for i, oriTextField := range oriTextFields {
			var chtTextField, engTextField string
			if len(chtTextFields) > i {
				chtTextField = chtTextFields[i]
			}
			if len(engTextFields) > i {
				engTextField = engTextFields[i]
			}
			result += fmt.Sprintf("text=%s&chtText=%s&engText=%s\n", oriTextField, chtTextField, engTextField)
		}
		c.WriteString(result)
	case "updatedb":
		if id == "lang" {
			text := c.Input.Query("text")
			chtText := c.Input.Query("chtText")
			engText := c.Input.Query("engText")
			controllers.InsertLang(text, chtText, engText)
			c.WriteString("OK")
		} else {
			c.WriteString("unsupport id:" + id)
		}
	default:
		c.WriteString("unsupport type:" + toolType)
	}
}
