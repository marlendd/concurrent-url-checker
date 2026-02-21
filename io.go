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
			if !strings.HasPrefix(u, "http") { u = "http://" + u }
			urls = append(urls, u)
		}
	}
	return urls, scanner.Err()
}

func renderReport(results []Result, totalTime time.Duration) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "TARGET URL\tSTATUS\tLATENCY\tNOTE")
	fmt.Fprintln(w, "---\t---\t---\t---")

	var success, failed int
	var totalLatency time.Duration

	for _, r := range results {
		status := fmt.Sprintf("%d", r.Status)
		note := "OK"
		
		if r.Err != nil {
			status = "ERR"
			note = "Offline/Timeout"
			failed++
		} else {
			success++
			totalLatency += r.Latency
		}

		fmt.Fprintf(w, "%s\t%s\t%v\t%s\n", 
			r.URL, status, r.Latency.Round(time.Millisecond), note)
	}
	w.Flush()

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Printf("STATISTICS:\n")
	fmt.Printf("Total Success:  %d\n", success)
	fmt.Printf("Total Failed:   %d\n", failed)
	
	if success > 0 {
		avg := totalLatency / time.Duration(success)
		fmt.Printf("Avg Latency:    %v\n", avg.Round(time.Millisecond))
	}
	
	fmt.Printf("Execution Time: %v\n", totalTime.Round(time.Millisecond))
	fmt.Println(strings.Repeat("=", 40))
}