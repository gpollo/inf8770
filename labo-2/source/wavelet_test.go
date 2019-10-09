package main

import (
	"testing"
)

func TestGetXLowPassFilter(t *testing.T) {
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

	result := data.GetXLowPassFilter()

	assert2DByteArrayEqual(t, result, expected)
}

func TestGetYLowPassFilter(t *testing.T) {
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

	result := data.GetYLowPassFilter()

	assert2DByteArrayEqual(t, result, expected)
}

func TestGetXHighPassFilter(t *testing.T) {
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

	result := data.GetXHighPassFilter()

	assert2DByteArrayEqual(t, result, expected)
}

func TestGetYHighPassFilter(t *testing.T) {
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

	result := data.GetYHighPassFilter()

	assert2DByteArrayEqual(t, result, expected)
}
