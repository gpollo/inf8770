package main

import (
	"testing"
)

const dim = 8

func assertEqual(t *testing.T, g, e int) {
	if g != e {
		t.Errorf("got=%d, expected=%d", g, e)
	}
}

func assertEqualLowPass(t *testing.T, g, e1, e2 byte) {
	assertEqual(t, int(g), int((e1+e2)/2))
}

func assertEqualHighPass(t *testing.T, g, e1, e2 byte) {
	assertEqual(t, int(g), int((e1-e2)/2))
}

func TestGetXLowPassFilter(t *testing.T) {
	data := NewImageData(dim, dim)
	for j := 0; j < dim; j++ {
		for i := 0; i < dim; i++ {
			data[j][i] = byte(dim*j + 1)
		}
	}

	result := data.GetXLowPassFilter()
	sizeX, sizeY := result.GetDimensions()

	assertEqual(t, sizeX, dim/2)
	assertEqual(t, sizeY, dim/1)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			assertEqualLowPass(t, result[j][i], data[j][2*i+0], data[j][2*i+1])
		}
	}
}

func TestGetYLowPassFilter(t *testing.T) {
	data := NewImageData(dim, dim)
	for j := 0; j < dim; j++ {
		for i := 0; i < dim; i++ {
			data[j][i] = byte(dim*j + 1)
		}
	}

	result := data.GetYLowPassFilter()
	sizeX, sizeY := result.GetDimensions()

	assertEqual(t, sizeX, dim/1)
	assertEqual(t, sizeY, dim/2)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			assertEqualLowPass(t, result[j][i], data[2*j+0][i], data[2*j+1][i])
		}
	}
}

func TestGetXHighPassFilter(t *testing.T) {
	data := NewImageData(dim, dim)
	for j := 0; j < dim; j++ {
		for i := 0; i < dim; i++ {
			data[j][i] = byte(dim*j + 1)
		}
	}

	result := data.GetXHighPassFilter()
	sizeX, sizeY := result.GetDimensions()

	assertEqual(t, sizeX, dim/2)
	assertEqual(t, sizeY, dim/1)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			assertEqualHighPass(t, result[j][i], data[j][2*i+0], data[j][2*i+1])
		}
	}
}

func TestGetYHighPassFilter(t *testing.T) {
	data := NewImageData(dim, dim)
	for j := 0; j < dim; j++ {
		for i := 0; i < dim; i++ {
			data[j][i] = byte(dim*j + 1)
		}
	}

	result := data.GetYHighPassFilter()
	sizeX, sizeY := result.GetDimensions()

	assertEqual(t, sizeX, dim/1)
	assertEqual(t, sizeY, dim/2)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			assertEqualHighPass(t, result[j][i], data[2*j+0][i], data[2*j+1][i])
		}
	}
}
