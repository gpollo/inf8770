package main

import (
	"math/rand"
	"testing"
)

func TestLZWCompressor(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	compressor := NewLZWCompressor()
	compressed := compressor.CompressLayer(data)
	decompressed := compressor.DecompressLayer(compressed)

	assert2DFloat32ArrayEqual(t, decompressed, data)
}

func TestLZWCompressorRandom(t *testing.T) {
	size := 1000
	data := NewImageData(size, size)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			data[j][i] = float32(rand.Int31n(256))
		}
	}

	compressor := NewLZWCompressor()
	compressed := compressor.CompressLayer(data)
	decompressed := compressor.DecompressLayer(compressed)

	assert2DFloat32ArrayEqual(t, decompressed, data)
}
