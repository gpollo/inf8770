package main

import (
	"testing"
)

func assert2DByteArrayEqual(t *testing.T, got, expected [][]byte) {
	if len(got) != len(expected) {
		t.Fatalf("array size differs: got=%d, expected=%d", len(got), len(expected))
	}

	if len(got) == 0 {
		return
	}

	if len(got[0]) != len(expected[0]) {
		t.Fatalf("array size differs: got=%d, expected=%d", len(got[0]), len(expected[0]))
	}

	for j := 0; j < len(got); j++ {
		for i := 0; i < len(got[0]); i++ {
			if got[j][i] == expected[j][i] {
				continue
			}

			t.Errorf("value at (%d, %d) differs: got=%d, expected=%d", i, j, got[j][i], expected[j][i])
		}
	}
}

func TestSubsampler410(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	yExpected := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvExpected := ImageData{
		{23, 38},
		{46, 32},
	}

	subsampler := Subsampler410{}
	y, u, v := subsampler.Subsample(data, data, data)

	assert2DByteArrayEqual(t, y, yExpected)
	assert2DByteArrayEqual(t, u, uvExpected)
	assert2DByteArrayEqual(t, v, uvExpected)
}

func TestSubsampler420(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	yExpected := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvExpected := ImageData{
		{25, 22, 45, 32},
		{52, 40, 27, 37},
	}

	subsampler := Subsampler420{}
	y, u, v := subsampler.Subsample(data, data, data)

	assert2DByteArrayEqual(t, y, yExpected)
	assert2DByteArrayEqual(t, u, uvExpected)
	assert2DByteArrayEqual(t, v, uvExpected)
}

func TestSubsampler422(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	yExpected := ImageData{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvExpected := ImageData{
		{10, 20, 30, 45},
		{40, 25, 60, 20},
		{90, 35, 40, 40},
		{15, 45, 15, 35},
	}

	subsampler := Subsampler422{}
	y, u, v := subsampler.Subsample(data, data, data)

	assert2DByteArrayEqual(t, y, yExpected)
	assert2DByteArrayEqual(t, u, uvExpected)
	assert2DByteArrayEqual(t, v, uvExpected)
}
