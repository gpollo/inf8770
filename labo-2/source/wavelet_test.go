package main

import (
	"math/rand"
	"testing"
)

func TestHaarGetXLowPassFilter(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	expected := ImageData{
		{10, 20, 30, 45},
		{40, 25, 60, 20},
		{90, 35, 40, 40},
		{15, 45, 15, 35},
	}
	result := wavelet.GetXLowPassFilter(data)
	assert2DFloat32ArrayEqual(t, result, expected)
}

func TestHaarGetYLowPassFilter(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	expected := ImageData{
		{20, 30, 15, 30, 55, 35, 30, 35},
		{50, 55, 20, 60, 15, 40, 20, 55},
	}
	result := wavelet.GetYLowPassFilter(data)
	assert2DFloat32ArrayEqual(t, result, expected)
}

func TestHaarGetXHighPassFilter(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	expected := ImageData{
		{0, 0, 0, -5},
		{-10, -15, 20, 0},
		{0, -25, -20, -30},
		{-5, -15, -5, -5},
	}
	result := wavelet.GetXHighPassFilter(data)
	assert2DFloat32ArrayEqual(t, result, expected)
}

func TestHaarGetYHighPassFilter(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	expected := ImageData{
		{-10, -20, 5, -10, -25, -5, 10, 15},
		{40, 35, -10, 0, 5, 20, -10, 15},
	}
	result := wavelet.GetYHighPassFilter(data)
	assert2DFloat32ArrayEqual(t, result, expected)
}

func TestHaarScaleX(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20},
		{30, 45, 10, 40},
		{40, 40, 10, 45},
		{10, 20, 30, 45},
	}

	expected := ImageData{
		{20, 0, 20, 0, 40, 0, 40, 0},
		{60, 0, 90, 0, 20, 0, 80, 0},
		{80, 0, 80, 0, 20, 0, 90, 0},
		{20, 0, 40, 0, 60, 0, 90, 0},
	}
	result := wavelet.ScaleX(data, data)
	assert2DFloat32ArrayEqual(t, result, expected)
}

func TestHaarScaleY(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20},
		{30, 45, 10, 40},
		{40, 40, 10, 45},
		{10, 20, 30, 45},
	}

	expected := ImageData{
		{20, 20, 40, 40},
		{00, 00, 00, 00},
		{60, 90, 20, 80},
		{00, 00, 00, 00},
		{80, 80, 20, 90},
		{00, 00, 00, 00},
		{20, 40, 60, 90},
		{00, 00, 00, 00},
	}
	result := wavelet.ScaleY(data, data)
	assert2DFloat32ArrayEqual(t, result, expected)
}

func TestHaarCopyIntoQuadrant(t *testing.T) {
	wavelet := HaarWavelet{}
	into := ImageData{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	from1 := ImageData{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	expected1 := ImageData{
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	wavelet.CopyIntoQuadrant(from1, into, 1)
	assert2DFloat32ArrayEqual(t, into, expected1)

	from2 := ImageData{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}
	expected2 := ImageData{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	wavelet.CopyIntoQuadrant(from2, into, 2)
	assert2DFloat32ArrayEqual(t, into, expected2)

	from3 := ImageData{
		{3, 3, 3, 3},
		{3, 3, 3, 3},
		{3, 3, 3, 3},
		{3, 3, 3, 3},
	}
	expected3 := ImageData{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{3, 3, 3, 3, 0, 0, 0, 0},
		{3, 3, 3, 3, 0, 0, 0, 0},
		{3, 3, 3, 3, 0, 0, 0, 0},
		{3, 3, 3, 3, 0, 0, 0, 0},
	}
	wavelet.CopyIntoQuadrant(from3, into, 3)
	assert2DFloat32ArrayEqual(t, into, expected3)

	from4 := ImageData{
		{4, 4, 4, 4},
		{4, 4, 4, 4},
		{4, 4, 4, 4},
		{4, 4, 4, 4},
	}
	expected4 := ImageData{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
	}
	wavelet.CopyIntoQuadrant(from4, into, 4)
	assert2DFloat32ArrayEqual(t, into, expected4)

	from5 := ImageData{
		{5, 5, 5, 5},
		{5, 5, 5, 5},
	}
	expected5 := ImageData{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
	}
	wavelet.CopyIntoQuadrant(from5, into, 4)
	assert2DFloat32ArrayEqual(t, into, expected5)

	from6 := ImageData{
		{6, 6},
		{6, 6},
		{6, 6},
		{6, 6},
	}
	expected6 := ImageData{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{3, 3, 6, 6, 4, 4, 4, 4},
		{3, 3, 6, 6, 4, 4, 4, 4},
		{3, 3, 6, 6, 4, 4, 4, 4},
		{3, 3, 6, 6, 4, 4, 4, 4},
	}
	wavelet.CopyIntoQuadrant(from6, into, 4)
	assert2DFloat32ArrayEqual(t, into, expected6)
}

func TestHaarWaveletTransform(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	wavelet1 := HaarWavelet{level: 1}
	expected1 := ImageData{
		{25.00, 22.50, 45.00, 32.50, -15.00, -2.50, -15.00, 12.50},
		{52.50, 40.00, 27.50, 37.50, 37.50, -5.00, 12.50, 2.50},
		{-5.00, -7.50, 10.00, -2.50, 5.00, 7.50, -10.00, -2.50},
		{-2.50, -20.00, -12.50, -17.50, 2.50, -5.00, -7.50, -12.50},
	}
	result1 := wavelet1.WaveletTransform(data)
	assert2DFloat32ArrayEqual(t, result1, expected1)

	wavelet2 := HaarWavelet{level: 2}
	expected2 := ImageData{
		{35.000, 35.625, -11.250, 3.125, -15.000, -2.500, -15.000, 12.500},
		{3.750, 0.625, -2.500, 5.625, 37.500, -5.000, 12.500, 2.500},
		{-5.000, -7.500, 10.000, -2.500, 5.000, 7.500, -10.000, -2.500},
		{-2.500, -20.000, -12.500, -17.500, 2.500, -5.00, -7.500, -12.500},
	}
	result2 := wavelet2.WaveletTransform(data)
	assert2DFloat32ArrayEqual(t, result2, expected2)
}

func TestHaarWavelet(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	wavelet1 := HaarWavelet{level: 1}
	transformed1 := wavelet1.WaveletTransform(data)
	inversed1 := wavelet1.WaveletInverse(transformed1)
	assert2DFloat32ArrayEqual(t, inversed1, data)

	wavelet2 := HaarWavelet{level: 2}
	transformed2 := wavelet2.WaveletTransform(data)
	inversed2 := wavelet2.WaveletInverse(transformed2)
	assert2DFloat32ArrayEqual(t, inversed2, data)
}

func TestHaarWaveletRandom(t *testing.T) {
	size := 1000
	data := NewImageData(size, size)
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			data[j][i] = float32(rand.Int31n(256))
		}
	}

	wavelet := HaarWavelet{level: 8}
	transformed := wavelet.WaveletTransform(data)
	inversed := wavelet.WaveletInverse(transformed)
	assert2DFloat32ArrayEqual(t, inversed, data)
}
