package quantifier

import (
	"jpeg2000/data"
	"jpeg2000/helper"
	"testing"
)

func TestMidThreadQuantifier(t *testing.T) {
	input := data.Layer{
		{10, 10, 20, 20, 30, -30, 35, 50},
		{-30, 50, 10, 40, -80, 40, -20, 20},
		{90, -90, 10, -60, 20, 60, 10, -70},
		{10, 20, -30, 60, 10, 20, 30, 40},
	}

	quantifier := MidThreadQuantifier{delta: 15}
	quantified := quantifier.QuantifierTransform(input)
	result := quantifier.QuantifierInverse(quantified)
	expected := data.Layer{
		{22.5, 22.5, 22.50, 22.5, 37.5, -22.5, 37.5, 52.5},
		{-22.5, 52.5, 22.5, 52.5, -67.5, 52.5, -7.5, 22.5},
		{97.5, -82.5, 22.5, -52.5, 22.5, 67.5, 22.5, -67.5},
		{22.5, 22.5, -22.5, 67.5, 22.5, 22.5, 37.5, 52.5},
	}
	helper.Assert2DFloat32ArrayAlmostEqual(t, result, expected, 0.001)
}
