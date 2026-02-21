package main

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL     string
	Status  int
	Latency time.Duration
	Err     error
}

func checkAll(urls []string, workerCount int) []Result {
	jobs := make(chan string, len(urls))
	resultsChan := make(chan Result, len(urls))
	var wg sync.WaitGroup

	// запуск пула воркеров
	for i := 0; i < workerCount; i++ {
		wg.Go(func() {
			client := &http.Client{Timeout: 10 * time.Second}
			for url := range jobs {
				start := time.Now()
				resp, err := client.Get(url)
				
				res := Result{URL: url, Latency: time.Since(start)}
				if err != nil {
					res.Err = err
				} else {
					res.Status = resp.StatusCode
					resp.Body.Close()
				}
				resultsChan <- res
			}
		})
		
	}

	// раздача задач
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
	close(resultsChan)

	var results []Result
	for r := range resultsChan {
		results = append(results, r)
	}
	return results
}