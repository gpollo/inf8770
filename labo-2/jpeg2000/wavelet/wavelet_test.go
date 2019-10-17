package wavelet

import (
	"jpeg2000/data"
	"jpeg2000/helper"
	"testing"
)

func TestScaleX(t *testing.T) {
	input := data.Layer{
		{10, 10, 20, 20},
		{30, 45, 10, 40},
		{40, 40, 10, 45},
		{10, 20, 30, 45},
	}

	expected := data.Layer{
		{20, 0, 20, 0, 40, 0, 40, 0},
		{60, 0, 90, 0, 20, 0, 80, 0},
		{80, 0, 80, 0, 20, 0, 90, 0},
		{20, 0, 40, 0, 60, 0, 90, 0},
	}
	result := scaleX(input, input)
	helper.Assert2DFloat32ArrayEqual(t, result, expected)
}

func TestScaleY(t *testing.T) {
	input := data.Layer{
		{10, 10, 20, 20},
		{30, 45, 10, 40},
		{40, 40, 10, 45},
		{10, 20, 30, 45},
	}

	expected := data.Layer{
		{20, 20, 40, 40},
		{00, 00, 00, 00},
		{60, 90, 20, 80},
		{00, 00, 00, 00},
		{80, 80, 20, 90},
		{00, 00, 00, 00},
		{20, 40, 60, 90},
		{00, 00, 00, 00},
	}
	result := scaleY(input, input)
	helper.Assert2DFloat32ArrayEqual(t, result, expected)
}

func TestCopyIntoQuadrant(t *testing.T) {
	into := data.Layer{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	from1 := data.Layer{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	expected1 := data.Layer{
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	copyIntoQuadrant(from1, into, 1)
	helper.Assert2DFloat32ArrayEqual(t, into, expected1)

	from2 := data.Layer{
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
		{2, 2, 2, 2},
	}
	expected2 := data.Layer{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	copyIntoQuadrant(from2, into, 2)
	helper.Assert2DFloat32ArrayEqual(t, into, expected2)

	from3 := data.Layer{
		{3, 3, 3, 3},
		{3, 3, 3, 3},
		{3, 3, 3, 3},
		{3, 3, 3, 3},
	}
	expected3 := data.Layer{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{3, 3, 3, 3, 0, 0, 0, 0},
		{3, 3, 3, 3, 0, 0, 0, 0},
		{3, 3, 3, 3, 0, 0, 0, 0},
		{3, 3, 3, 3, 0, 0, 0, 0},
	}
	copyIntoQuadrant(from3, into, 3)
	helper.Assert2DFloat32ArrayEqual(t, into, expected3)

	from4 := data.Layer{
		{4, 4, 4, 4},
		{4, 4, 4, 4},
		{4, 4, 4, 4},
		{4, 4, 4, 4},
	}
	expected4 := data.Layer{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
	}
	copyIntoQuadrant(from4, into, 4)
	helper.Assert2DFloat32ArrayEqual(t, into, expected4)

	from5 := data.Layer{
		{5, 5, 5, 5},
		{5, 5, 5, 5},
	}
	expected5 := data.Layer{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
		{3, 3, 3, 3, 4, 4, 4, 4},
	}
	copyIntoQuadrant(from5, into, 4)
	helper.Assert2DFloat32ArrayEqual(t, into, expected5)

	from6 := data.Layer{
		{6, 6},
		{6, 6},
		{6, 6},
		{6, 6},
	}
	expected6 := data.Layer{
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 1, 1, 1, 1},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{2, 2, 2, 2, 5, 5, 5, 5},
		{3, 3, 6, 6, 4, 4, 4, 4},
		{3, 3, 6, 6, 4, 4, 4, 4},
		{3, 3, 6, 6, 4, 4, 4, 4},
		{3, 3, 6, 6, 4, 4, 4, 4},
	}
	copyIntoQuadrant(from6, into, 4)
	helper.Assert2DFloat32ArrayEqual(t, into, expected6)
}
