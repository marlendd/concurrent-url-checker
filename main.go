package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := flag.String("f", "sites.txt", "path to urls file")
	workers := flag.Int("w", 20, "number of workers")
	flag.Parse()

	urls, err := loadURLs(*filePath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	start := time.Now() 
	fmt.Printf("Starting check of %d URLs with %d workers...\n\n", len(urls), *workers)
	
	results := checkAll(urls, *workers)
	totalTime := time.Since(start)

	// Передаем результаты и общее время в отчет
	renderReport(results, totalTime)
}