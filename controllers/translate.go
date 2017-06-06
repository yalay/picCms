package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"conf"
)

const apiKey = "trnsl.1.1.20170604T030854Z.e1cac9fc98790241.b0defc0c4b712d7bde91e5d2a7621534cfaf3b23"

// {"code":200,"lang":"zh-en","text":["The body bare min Ni Mily ribbon bundle creative fun kit figure"]}
type rspMsg struct {
	Code int
	Lang string
	Text []string
}

func Translate(text, langType string) string {
	if langType == conf.KlangTypeCn {
		return text
	}

	cacheKey := MakeCacheKey(KcachePrefixLang, text, langType)
	if cacheData, err := CACHE.Get(cacheKey); err == nil {
		return cacheData.(string)
	}else {
		var transText string
		textFields := strings.Split(text, ",")
		if len(textFields) == 1 {
			transText = GetLang(text, langType)
		} else {
			var engTexts = make([]string, 0, len(textFields))
			for _, textField := range textFields {
				engTextField := GetLang(textField, langType)
				if engTextField == "" {
					continue
				}
				engTexts = append(engTexts, engTextField)
			}
			transText =strings.Join(engTexts, ",")
		}
		if transText == "" {
			return text
		}

		CACHE.Set(cacheKey, transText)
		return transText
	}
}

// 转换到繁体
func TranslateToCht(text string) (string, error) {
	baseUrl := "http://opencc.byvoid.com/convert"
	queryParam := url.Values{}
	queryParam.Set("text", text)
	queryParam.Set("config", "s2twp.json")
	queryParam.Set("precise", "0")
	resp, err := http.PostForm(baseUrl, queryParam)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// 转换到英文
func TranslateToEng(text string) (string, error) {
	baseUrl := "https://translate.yandex.net/api/v1.5/tr.json/translate?"
	queryParam := url.Values{}
	queryParam.Set("key", apiKey)
	queryParam.Set("text", text)
	queryParam.Set("lang", "zh-en")

	resp, err := http.Get(baseUrl+queryParam.Encode())
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
