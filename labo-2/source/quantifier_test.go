package main

import (
	"testing"
)

func TestDeadZoneQuantifier(t *testing.T) {
	data := ImageData{
		{10, 10, 20, 20, 30, -30, 35, 50},
		{-30, 50, 10, 40, -80, 40, -20, 20},
		{90, -90, 10, -60, 20, 60, 10, -70},
		{10, 20, -30, 60, 10, 20, 30, 40},
	}

	quantifier := DeadZoneQuantifier{
		width:  20,
		delta:  15,
		offset: 0,
	}
	quantified := quantifier.QuantifierTransform(data)
	result := quantifier.QuantifierInverse(quantified)
	expected := ImageData{
		{5, 5, 20, 20, 20, -20, 35, 50},
		{-20, 50, 5, 35, -80, 35, -20, 20},
		{80, -80, 5, -50, 20, 50, 5, -65},
		{5, 20, -20, 50, 5, 20, 20, 35},
	}
	assert2DFloat32ArrayEqual(t, result, expected)
}
