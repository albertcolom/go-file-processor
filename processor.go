package main

import (
	"bufio"
	"io"
	"log"
	"sync"
)

func processFile(file io.Reader, numWorkers int, chunkSize int) {
	scanner := bufio.NewScanner(file)
	payloadChan := make(chan *Chunk, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for chunk := range payloadChan {
				for _, payload := range chunk.payloads {
					// Actual processing logic goes here
					// fmt.Println(string(payload.Data))

					releasePayload(payload)
				}

				releaseChunk(chunk)
			}
		}(i)
	}

	go func() {
		defer close(payloadChan)
		for {
			chunk := getChunk()
			for len(chunk.payloads) < chunkSize && scanner.Scan() {
				payload := getPayload()

				payload.Data = payload.Data[:0]
				payload.Data = append(payload.Data[:0], scanner.Bytes()...)

				chunk.payloads = append(chunk.payloads, payload)
			}

			if len(chunk.payloads) > 0 {
				payloadChan <- chunk
			} else {
				releaseChunk(chunk)
				break
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Scanner error: %v", err)
		}
	}()

	wg.Wait()
}
