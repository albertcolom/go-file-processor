package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	filePath := flag.String("file", "", "Path to CSV file to process (required)")
	chunkSize := flag.Int("chunk", 1000, "Number of lines per chunk (default: 1000)")
	workers := flag.Int("workers", runtime.NumCPU(), "Number of concurrent workers (default: number of CPUs)")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Missing required --file parameter")
	}

	startTime := time.Now()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file stats: %v", err)
	}

	fileSize := stat.Size()
	log.Printf("Processing file: %s (%.2f MB)", *filePath, float64(fileSize)/1024/1024)
	log.Printf("Number of workers: %d", *workers)
	log.Printf("Chunks size: %d", *chunkSize)

	processFile(file, *workers, *chunkSize)

	printMemStats()
	log.Printf("Processing completed in %v", time.Since(startTime))
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("Memory stats: Alloc=%.2fMB, TotalAlloc=%.2fMB, Sys=%.2fMB, NumGC=%d",
		float64(m.Alloc)/1024/1024,
		float64(m.TotalAlloc)/1024/1024,
		float64(m.Sys)/1024/1024,
		m.NumGC,
	)
}
