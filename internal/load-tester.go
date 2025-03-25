package internal

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LoadTester struct {
	URL          string
	Requests     int
	Concurrency  int
	StatusCounts map[int]int
	Mutex        sync.Mutex
}

func NewLoadTester(url string, requests, concurrency int) *LoadTester {
	return &LoadTester{
		URL:          url,
		Requests:     requests,
		Concurrency:  concurrency,
		StatusCounts: make(map[int]int),
	}
}

func (lt *LoadTester) performRequest(wg *sync.WaitGroup, channelConcorrence chan struct{}) {
	defer wg.Done()
	defer func() { <-channelConcorrence }()

	resp, err := http.Get(lt.URL)
	if err != nil {
		fmt.Println("Erro na requisição:", err)
		return
	}
	defer resp.Body.Close()

	lt.Mutex.Lock()
	lt.StatusCounts[resp.StatusCode]++
	lt.Mutex.Unlock()
}

func (lt *LoadTester) Run() time.Duration {
	channelConcorrence := make(chan struct{}, lt.Concurrency)
	var wg sync.WaitGroup

	startTime := time.Now()
	for i := 0; i < lt.Requests; i++ {
		channelConcorrence <- struct{}{}
		wg.Add(1)
		go lt.performRequest(&wg, channelConcorrence)
	}

	wg.Wait()
	close(channelConcorrence)
	return time.Since(startTime)
}

func (lt *LoadTester) Report(elapsedTime time.Duration) {
	fmt.Println("\n--- Relatório ---")
	fmt.Printf("Tempo total: %s\n", elapsedTime)
	fmt.Printf("Total de requests: %d\n", lt.Requests)
	for status, count := range lt.StatusCounts {
		fmt.Printf("Status %d: %d\n", status, count)
	}
}
