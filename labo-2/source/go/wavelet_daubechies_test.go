package main

import (
	"math/rand"
	"testing"
)

func TestDaubechiesWavelet(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	wavelet1 := NewDaubechiesWavelet()
	wavelet1.SetLevel(1)
	transformed1 := wavelet1.WaveletTransform(data)
	inversed1 := wavelet1.WaveletInverse(transformed1)
	assert2DFloat32ArrayAlmostEqual(t, inversed1, data, 0.0001)

	wavelet2 := NewDaubechiesWavelet()
	wavelet2.SetLevel(2)
	transformed2 := wavelet2.WaveletTransform(data)
	inversed2 := wavelet2.WaveletInverse(transformed2)
	assert2DFloat32ArrayAlmostEqual(t, inversed2, data, 0.0001)
}

func TestDaubechiesWaveletRandom(t *testing.T) {
	size := 200
	data := NewImageData(size, size)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			data[j][i] = float32(rand.Int31n(256))
		}
	}

	wavelet := NewDaubechiesWavelet()
	wavelet.SetLevel(3)
	transformed := wavelet.WaveletTransform(data)
	inversed := wavelet.WaveletInverse(transformed)
	assert2DFloat32ArrayAlmostEqual(t, inversed, data, 0.0001)
}
