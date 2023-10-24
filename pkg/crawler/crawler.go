package crawler

import (
	"github.com/gocolly/colly"
	"time"
)

const MaxRetries = 3

func Crawl(url string) (string, error) {
	c := colly.NewCollector()

	var pageContent string
	var err error
	retries := 0

	c.OnResponse(func(r *colly.Response) {
		pageContent = string(r.Body)
	})

	for retries < MaxRetries {
		err = c.Visit(url)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		retries++
	}

	if err != nil {
		return "", err
	}

	c.Wait()

	return pageContent, nil
}
