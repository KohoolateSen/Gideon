package utils

import (
	"strings"
	"time"

	"Gideon/internal/utils/httpproxy"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
)

func Scrape(tag string, url string, charlimit int) string {
	var texts []string

	c := colly.NewCollector()

	c.WithTransport(httpproxy.GlobalHTTPClient.Transport)

	c.OnHTML(tag, func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		if text != "" && len(strings.Join(texts, " ")+text) <= charlimit {
			texts = append(texts, text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Infof("visiting: %s", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		logrus.Warn("request URL:", r.Request.URL, "failed with response: ", r, "\nError:", err)
	})

	c.SetRequestTimeout(30 * time.Second)

	c.Visit(url)

	joinedText := strings.Join(texts, " ")

	return joinedText
}
