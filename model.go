package main

import "sync"

type Payload struct {
	Data []byte
}

type Chunk struct {
	payloads []*Payload
}

var payloadPool = sync.Pool{
	New: func() interface{} {
		return &Payload{}
	},
}

var chunkPool = sync.Pool{
	New: func() interface{} {
		return &Chunk{
			payloads: make([]*Payload, 0),
		}
	},
}

func getPayload() *Payload {
	return payloadPool.Get().(*Payload)
}

func releasePayload(payload *Payload) {
	payload.Data = payload.Data[:0]
	payloadPool.Put(payload)
}

func getChunk() *Chunk {
	return chunkPool.Get().(*Chunk)
}

func releaseChunk(chunk *Chunk) {
	chunk.payloads = chunk.payloads[:0]
	chunkPool.Put(chunk)
}
