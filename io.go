package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func loadURLs(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		u := strings.TrimSpace(scanner.Text())
		if u != "" {
			if !strings.HasPrefix(u, "http") {
				u = "http://" + u
			}
			urls = append(urls, u)
		}
	}
	return urls, scanner.Err()
}

func renderReport(results []Result) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "TARGET URL\tSTATUS\tLATENCY\tNOTE")
	fmt.Fprintln(w, "---\t---\t---\t---")

	for _, r := range results {
		status := fmt.Sprintf("%d", r.Status)
		note := "OK"
		
		if r.Err != nil {
			status = "ERR"
			note = r.Err.Error()
			if len(note) > 30 { note = note[:30] + "..." } // Обрезаем длинные ошибки
		} else if r.Status != 200 {
			note = "Check site"
		}

		fmt.Fprintf(w, "%s\t%s\t%v\t%s\n", 
			r.URL, status, r.Latency.Round(time.Millisecond), note)
	}
	w.Flush()
}