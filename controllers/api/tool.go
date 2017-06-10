package api

import (
	"controllers"
	"sync"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego/context"
)

func ToolHandler(c *context.Context) {
	toolType := c.Input.Param(":type")
	switch toolType {
	case "translate":
		articleId := Atoi32(c.Input.Param(":id"))
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
		for i, oriTextField := range oriTextFields {
			var chtTextField, engTextField string
			if len(chtTextFields) > i {
				chtTextField = chtTextFields[i]
			}
			if len(engTextFields) > i {
				engTextField = engTextFields[i]
			}
			//controllers.InsertLang(oriTextField, chtTextField, engTextField)
			log.Println(oriTextField + "->" + chtTextField + ";" + engTextField)
		}
		c.WriteString("OK")
	default:
		c.WriteString("unsupport type.")
	}
}
