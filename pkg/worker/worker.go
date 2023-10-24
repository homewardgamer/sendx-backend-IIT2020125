package worker

import (
	"github.com/TheRealSibasishBehera/go-web-crawler/pkg/crawler"
	"github.com/TheRealSibasishBehera/go-web-crawler/pkg/storage"
	"sync"
	"time"
)

type Job struct {
	URL        string
	IsPaying   bool
	ResultChan chan string
	ErrorChan  chan error
}

var mu sync.Mutex
var jobQueue = make(chan Job, 100)

var payingWorkerCount = 5
var freeWorkerCount = 2
var rateLimit = 100

func SetPayingWorkerCount(count int) {
	mu.Lock()
	defer mu.Unlock()
	payingWorkerCount = count
}

func SetFreeWorkerCount(count int) {
	mu.Lock()
	defer mu.Unlock()
	freeWorkerCount = count
}

func SetRateLimit(pages int) {
	mu.Lock()
	defer mu.Unlock()
	rateLimit = pages
}

func StartWorkers() {
	for i := 0; i < payingWorkerCount; i++ {
		go worker(true)
	}

	for i := 0; i < freeWorkerCount; i++ {
		go worker(false)
	}
}

func worker(isPaying bool) {
	ticker := time.NewTicker(time.Hour / time.Duration(rateLimit))

	for {
		select {
		case job := <-jobQueue:
			if job.IsPaying == isPaying {
				<-ticker.C
				result, err := crawl(job.URL)
				if err != nil {
					job.ErrorChan <- err
				} else {
					job.ResultChan <- result
				}
			}
		}
	}
}

func crawl(url string) (string, error) {
	content, err := crawler.Crawl(url)

	if err != nil {
		return "", err
	}

	err = storage.SavePage(url, content)

	if err != nil {
		return "", err
	}

	return content, nil
}

func QueueJob(url string, isPaying bool) (chan string, chan error) {
	content, err := storage.GetPage(url)

	if err == nil {
		resultChan := make(chan string, 1)
		errorChan := make(chan error, 1)

		resultChan <- content

		close(resultChan)
		close(errorChan)

		return resultChan, errorChan
	}

	job := Job{
		URL:        url,
		IsPaying:   isPaying,
		ResultChan: make(chan string),
		ErrorChan:  make(chan error),
	}
	jobQueue <- job

	return job.ResultChan, job.ErrorChan
}
