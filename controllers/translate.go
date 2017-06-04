package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const apiKey = "trnsl.1.1.20170604T030854Z.e1cac9fc98790241.b0defc0c4b712d7bde91e5d2a7621534cfaf3b23"

// {"code":200,"lang":"zh-en","text":["The body bare min Ni Mily ribbon bundle creative fun kit figure"]}
type rspMsg struct {
	Code int
	Lang string
	Text []string
}

func Translate(text, langType string) string {
	if langType != "en" {
		return text
	}

	cacheKey := MakeCacheKey(KcachePrefixLang, text, langType)
	if cacheData, err := CACHE.Get(cacheKey); err == nil {
		return cacheData.(string)
	}else {
		var engText string
		textFields := strings.Split(text, ",")
		if len(textFields) == 1 {
			engText = GetEngLang(text, langType)
		} else {
			var engTexts = make([]string, 0, len(textFields))
			for _, textField := range textFields {
				engTextField := GetEngLang(textField, langType)
				if engTextField == "" {
					continue
				}
				engTexts = append(engTexts, engTextField)
			}
			engText =strings.Join(engTexts, ",")
		}
		if engText == "" {
			return text
		}

		CACHE.Set(cacheKey, engText)
		return engText
	}
}

func TranslateToEng(text string) (string, error) {
	baseUrl := "https://translate.yandex.net/api/v1.5/tr.json/translate?"
	queryParam := url.Values{}
	queryParam.Set("key", apiKey)
	queryParam.Set("text", text)
	queryParam.Set("lang", "zh-en")
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+queryParam.Encode(), nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var msg = &rspMsg{}
	json.Unmarshal(body, msg)
	if msg.Code != 200 || len(msg.Text) == 0 {
		return "", fmt.Errorf("translate failed.\n")
	}

	return msg.Text[0], nil
}
