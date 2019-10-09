package main

import (
	"testing"
)

func TestGetXLowPassFilter(t *testing.T) {
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
	assert2DByteArrayEqual(t, result, expected)
}

func TestGetYLowPassFilter(t *testing.T) {
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
	assert2DByteArrayEqual(t, result, expected)
}

func TestGetXHighPassFilter(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	expected := ImageData{
		{0, 0, 0, 0},
		{0, 0, 20, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	result := wavelet.GetXHighPassFilter(data)
	assert2DByteArrayEqual(t, result, expected)
}

func TestGetYHighPassFilter(t *testing.T) {
	wavelet := HaarWavelet{}
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	expected := ImageData{
		{0, 0, 5, 0, 0, 0, 10, 15},
		{40, 35, 0, 0, 5, 20, 0, 15},
	}
	result := wavelet.GetYHighPassFilter(data)
	assert2DByteArrayEqual(t, result, expected)
}

func TestCopyIntoQuadrant(t *testing.T) {
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
	assert2DByteArrayEqual(t, into, expected1)

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
	assert2DByteArrayEqual(t, into, expected2)

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
	assert2DByteArrayEqual(t, into, expected3)

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
	assert2DByteArrayEqual(t, into, expected4)

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
	assert2DByteArrayEqual(t, into, expected5)

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
	assert2DByteArrayEqual(t, into, expected6)
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
		{25, 22, 45, 32, 0, 0, 0, 12},
		{52, 40, 27, 37, 37, 0, 12, 2},
		{0, 0, 10, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	result1 := wavelet1.WaveletTransform(data)
	assert2DByteArrayEqual(t, result1, expected1)

	wavelet2 := HaarWavelet{level: 2}
	expected2 := ImageData{
		{34, 35, 0, 3, 0, 0, 0, 12},
		{3, 3, 0, 3, 37, 0, 12, 2},
		{0, 0, 10, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	result2 := wavelet2.WaveletTransform(data)
	assert2DByteArrayEqual(t, result2, expected2)
}
