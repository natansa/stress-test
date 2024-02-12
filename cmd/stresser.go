package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	TotalDuration      time.Duration
	TotalRequests      int
	SuccessResponses   int
	StatusDistribution map[int]int
}

func Start(url string, totalRequests, concurrency int) RunEFunc {
	startTime := time.Now()
	var wg sync.WaitGroup
	results := make(chan int, totalRequests)
	concurrentGoroutines := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		concurrentGoroutines <- struct{}{}
		wg.Add(1)

		go func() {
			defer wg.Done()
			response, err := http.Get(url)
			if err == nil {
				results <- response.StatusCode
				response.Body.Close()
			} else {
				results <- 0
			}
			<-concurrentGoroutines
		}()
	}

	wg.Wait()
	close(results)

	result := processResults(results, startTime, totalRequests)
	printResult(result)
	return nil
}

func processResults(results chan int, startTime time.Time, totalRequests int) Result {
	statusDistribution := make(map[int]int)
	successResponses := 0
	for statusCode := range results {
		if statusCode == http.StatusOK {
			successResponses++
		}
		statusDistribution[statusCode]++
	}
	totalDuration := time.Since(startTime)

	return Result{
		TotalDuration:      totalDuration,
		TotalRequests:      totalRequests,
		SuccessResponses:   successResponses,
		StatusDistribution: statusDistribution,
	}
}

func printResult(result Result) {
	fmt.Printf("Total Time Spent: %v\n", result.TotalDuration)
	fmt.Printf("Total Requests Made: %d\n", result.TotalRequests)
	fmt.Printf("Success Responses: %d\n", result.SuccessResponses)
	fmt.Println("HTTP Status Codes Distribution:")
	for status, count := range result.StatusDistribution {
		fmt.Printf("  %d: %d\n", status, count)
	}
}
