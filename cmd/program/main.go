package main

import (
	"flag"
	"fmt"

	"github.com/brunoofgod/goexpert-lesson-8/internal"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL é obrigatória")
		return
	}

	loadTester := internal.NewLoadTester(*url, *requests, *concurrency)
	elapsedTime := loadTester.Run()
	loadTester.Report(elapsedTime)
}
