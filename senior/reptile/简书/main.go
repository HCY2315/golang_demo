package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

// one
// func main() {
// 	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))
// 	//文章列表
// 	c.OnHTML("ul[class='note-list']", func(e *colly.HTMLElement) {
// 		//列表中每一项
// 		e.ForEach("li", func(i int, item *colly.HTMLElement) {
// 			//文章链接
// 			href := item.ChildAttr("div[class='content'] > a[class='title']", "href")
// 			//文章标题
// 			title := item.ChildText("div[class='content'] > a[class='title']")
// 			//文章摘要
// 			summary := item.ChildText("div[class='content'] > p[class='abstract']")
// 			fmt.Println(title, href)
// 			fmt.Println(summary)
// 			fmt.Println()
// 		})
// 	})

// 	err := c.Visit("https://www.jianshu.com")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

// two
type Hot struct {
	Href     string `selector:"a[class=title]" attr:"href"`
	Title    string `selector:"a[class=title]"`
	Summary  string `selector:"p[class=abstract]"`
	Nickname string `selector:"a[class=nickname]"`
}

func main() {
	hots := make([]*Hot, 0)
	// 创建 collector
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))

	// 事件监听，通过 callback 执行事件处理
	c.OnHTML("ul[class='note-list'],li", func(e *colly.HTMLElement) {
		hot := new(Hot)

		err := e.Unmarshal(hot)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		if hot.Summary != "" || hot.Href != "" || hot.Title != "" {
			hots = append(hots, hot)
		}
	})

	// 访问 AllowedDomains 中的链接
	err := c.Visit("https://www.jianshu.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, h := range hots {
		fmt.Println(h.Href + " || " + h.Title + " || " + h.Summary + " || " + h.Nickname)
	}
}
