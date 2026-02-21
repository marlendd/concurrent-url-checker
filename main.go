package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("f", "sites.txt", "файл со списком URL")
	workers := flag.Int("w", 15, "количество воркеров")
	flag.Parse()

	// 1. чтение URL
	urls, err := loadURLs(*filePath)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		os.Exit(1)
	}

	// 2. проверка
	fmt.Printf("Проверяем %d сайтов в %d потоков...\n\n", len(urls), *workers)
	results := checkAll(urls, *workers)

	// 3. вывод таблицы
	renderReport(results)
}