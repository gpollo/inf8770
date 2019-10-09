package main

import (
	"testing"
)

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
