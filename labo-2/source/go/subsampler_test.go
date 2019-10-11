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
		{23.75, 38.75},
		{46.25, 32.50},
	}

	subsampler := Subsampler410{}
	y, u, v := subsampler.Subsample(data, data, data)

	assert2DFloat32ArrayEqual(t, y, yExpected)
	assert2DFloat32ArrayEqual(t, u, uvExpected)
	assert2DFloat32ArrayEqual(t, v, uvExpected)
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
		{25.00, 22.50, 45.00, 32.50},
		{52.50, 40.00, 27.50, 37.50},
	}

	subsampler := Subsampler420{}
	y, u, v := subsampler.Subsample(data, data, data)

	assert2DFloat32ArrayEqual(t, y, yExpected)
	assert2DFloat32ArrayEqual(t, u, uvExpected)
	assert2DFloat32ArrayEqual(t, v, uvExpected)
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
		{10.00, 20.00, 30.00, 45.00},
		{40.00, 25.00, 60.00, 20.00},
		{90.00, 35.00, 40.00, 40.00},
		{15.00, 45.00, 15.00, 35.00},
	}

	subsampler := Subsampler422{}
	y, u, v := subsampler.Subsample(data, data, data)

	assert2DFloat32ArrayEqual(t, y, yExpected)
	assert2DFloat32ArrayEqual(t, u, uvExpected)
	assert2DFloat32ArrayEqual(t, v, uvExpected)
}
