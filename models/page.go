package models

import (
	"bytes"
	"strconv"
)

/*
<a class="page-numbers current">1</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-2.html">2</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-3.html">3</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-4.html">4</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-5.html">5</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-6.html">6</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-7.html">7</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-8.html">8</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-9.html">9</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-10.html">10</a>
<a class='page-numbers' href="http://www.xxxx.com/article-59-11.html">11</a>
<a class='page-numbers'>...</a> <a class='page-numbers' href="http://www.xxxx.com/article-59-19.html">19</a>
<a class="next page-numbers" href="http://www.xxxx.com/article-59-2.html"><i class="fa fa-chevron-right"></i></a>
*/

// 用于处理分页
type Page struct {
	TotalNum  int
	CurNum    int
	SizeNum   int // 最多显示多少页码
	UrlPrefix string
	UrlSuffix string
}

func (p *Page) Html() string {
	// 固定显示上一页链接和第一页链接，除非就是第一页
	buff := bytes.Buffer{}
	if p.CurNum == 1 {
		buff.WriteString(`<a class="current page-numbers" href="` + p.UrlPrefix + p.UrlSuffix +
			`">1</a>`)
	} else {
		if p.CurNum == 2 {
			buff.WriteString(`<a class="prev page-numbers" href="` +
				p.UrlPrefix + p.UrlSuffix + `"><i class="fa fa-chevron-left"></i></a>`)
		} else {
			buff.WriteString(`<a class="prev page-numbers" href="` + p.UrlPrefix + "-" +
				strconv.Itoa(p.CurNum-1) + p.UrlSuffix + `"><i class="fa fa-chevron-left"></i></a>`)
		}
		buff.WriteString(`<a class="page-numbers" href="` + p.UrlPrefix + p.UrlSuffix +
			`">1</a>`)
	}

	// 分页列表
	halfSizeNum := p.SizeNum / 2
	var startNum, endNum int
	if p.TotalNum <= p.SizeNum {
		startNum = 2
		endNum = p.TotalNum - 1
	} else {
		if p.CurNum-halfSizeNum > 1 {
			startNum = p.CurNum - halfSizeNum
		} else {
			startNum = 2
		}

		if startNum + p.SizeNum >= p.TotalNum {
			endNum = p.TotalNum - 1
		} else {
			endNum = startNum + p.SizeNum
		}
	}

	if startNum > 2 {
		buff.WriteString(`<a class="page-numbers">...</a>`)
	}

	for i := startNum; i <= endNum; i++ {
		pageText := strconv.Itoa(i)
		if i == p.CurNum {
			buff.WriteString(`<a class="current page-numbers">` + pageText + `</a>`)
		} else {
			buff.WriteString(`<a class="page-numbers" href="` +
				p.UrlPrefix + "-" + pageText + p.UrlSuffix + `">` + pageText + `</a>`)
		}
	}

	if endNum != p.TotalNum-1 {
		buff.WriteString(`<a class="page-numbers">...</a>`)
	}


	// 固定显示最后一页链接和下一页链接，除非就是最后一页
	if p.CurNum == p.TotalNum {
		buff.WriteString(`<a class="current page-numbers" href="` + p.UrlPrefix + "-" +
			strconv.Itoa(p.TotalNum) + p.UrlSuffix + `">` + strconv.Itoa(p.TotalNum) + `</a>`)
	} else {
		buff.WriteString(`<a class="page-numbers" href="` + p.UrlPrefix + "-" +
			strconv.Itoa(p.TotalNum) + p.UrlSuffix + `">` + strconv.Itoa(p.TotalNum) + `</a>`)
		buff.WriteString(`<a class="next page-numbers" href="` + p.UrlPrefix + "-" +
			strconv.Itoa(p.CurNum+1) + p.UrlSuffix + `">` + `<i class="fa fa-chevron-right"></i></a>`)
	}

	return buff.String()
}

func (p *Page) PreUrl() string {
	if p.CurNum == 1 {
		return "#"
	}

	if p.CurNum == 2 {
		return p.UrlPrefix + p.UrlSuffix
	} else {
		return p.UrlPrefix + "-" + strconv.Itoa(p.CurNum-1) + p.UrlSuffix
	}
}

func (p *Page) NextUrl() string {
	if p.CurNum == p.TotalNum {
		return "#"
	}

	return p.UrlPrefix + "-" + strconv.Itoa(p.CurNum+1) + p.UrlSuffix
}
