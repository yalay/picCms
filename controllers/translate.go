package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"conf"
	"regexp"
)

const apiKey = "trnsl.1.1.20170604T030854Z.e1cac9fc98790241.b0defc0c4b712d7bde91e5d2a7621534cfaf3b23"

var wordExp, _  = regexp.Compile(`[^\pP\s]+`)

// {"code":200,"lang":"zh-en","text":["The body bare min Ni Mily ribbon bundle creative fun kit figure"]}
type rspMsg struct {
	Code int
	Lang string
	Text []string
}

// 支持长句子，可以包含各种标点
func TranslateLongText(longText, langType string) string {
	return wordExp.ReplaceAllStringFunc(longText, func(text string) string {
		return Translate(text, langType)
	})
}

// 只支持无标点符号或者只有","的语句
func Translate(text, langType string) string {
	if langType == conf.KlangTypeCn {
		return text
	}

	var transText string
	textFields := strings.Split(text, ",")
	if len(textFields) == 1 {
		transText = GetLang(text, langType)
	} else {
		transTexts := BatchGetLang(textFields, langType)
		transText =strings.Join(transTexts, ",")
	}
	if transText == "" {
		return text
	}

	return transText
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
