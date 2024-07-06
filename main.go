package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&requests, "requests", 0, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas")
}

type Result struct {
	StatusCode int
	Duration   time.Duration
}

func worker(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)
	if err != nil {
		results <- Result{StatusCode: 0, Duration: duration}
		return
	}
	defer resp.Body.Close()
	results <- Result{StatusCode: resp.StatusCode, Duration: duration}
}

func main() {
	flag.Parse()

	if url == "" || requests == 0 || concurrency == 0 {
		flag.Usage()
		return
	}

	results := make(chan Result, requests)
	var wg sync.WaitGroup

	startTime := time.Now()
	for i := 0; i < requests; i++ {
		if i%concurrency == 0 {
			wg.Wait()
			wg = sync.WaitGroup{}
		}
		wg.Add(1)
		go worker(url, &wg, results)
	}
	wg.Wait()
	close(results)
	totalTime := time.Since(startTime)

	statusCount := make(map[int]int)
	totalRequests := 0
	for result := range results {
		totalRequests++
		statusCount[result.StatusCode]++
	}

	fmt.Printf("Tempo total gasto: %v\n", totalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", statusCount[http.StatusOK])
	fmt.Printf("Distribuição de outros códigos de status HTTP:\n")
	for status, count := range statusCount {
		if status != http.StatusOK {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
}
