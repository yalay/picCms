package api

import (
	"conf"
	"strings"
	"time"
)

const (
	kTimeLayout  = "2006-01-02 15:04:05"
	kTimeLayout2 = "2006-01-02"
)

func TimeFormat(ts int64) string {
	curTime := time.Unix(ts, 0)
	return curTime.Format(kTimeLayout)
}

func TimeFormat2(ts int64) string {
	curTime := time.Unix(ts, 0)
	return curTime.Format(kTimeLayout2)
}

func GetLang(headerLang string) string {
	if headerLang == "" {
		return conf.KlangTypeCn
	}

	lowerLang := strings.ToLower(headerLang)
	switch {
	case strings.HasPrefix(lowerLang, "en"):
		return conf.KlangTypeEn
	case strings.HasPrefix(lowerLang, "zh-tw"):
		return conf.KlangTypeTW
	case strings.HasPrefix(lowerLang, "zh-hk") ||
		strings.HasPrefix(lowerLang, "zh-mo") ||
		strings.HasPrefix(lowerLang, "zh-hant"):
		return conf.KlangTypeHK
	default:
		if strings.Contains(lowerLang, "zh") {
			return conf.KlangTypeCn
		} else {
			return conf.KlangTypeEn
		}
	}
}
