package subsampler

import (
	"jpeg2000/data"
	"jpeg2000/helper"
	"testing"
)

func TestSubsample410(t *testing.T) {
	yuvData := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	yExpected := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvExpected := data.Layer{
		{23.75, 38.75},
		{46.25, 32.50},
	}

	subsampler := Subsampler410{}
	y, u, v := subsampler.Subsample(yuvData, yuvData, yuvData)

	helper.Assert2DFloat32ArrayEqual(t, y, yExpected)
	helper.Assert2DFloat32ArrayEqual(t, u, uvExpected)
	helper.Assert2DFloat32ArrayEqual(t, v, uvExpected)
}

func TestSupersample410(t *testing.T) {
	yData := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvData := data.Layer{
		{23, 38},
		{46, 32},
	}

	uvExpected := data.Layer{
		{23, 23, 23, 23, 38, 38, 38, 38},
		{23, 23, 23, 23, 38, 38, 38, 38},
		{46, 46, 46, 46, 32, 32, 32, 32},
		{46, 46, 46, 46, 32, 32, 32, 32},
	}

	subsampler := Subsampler410{}
	y, u, v := subsampler.Supersample(yData, uvData, uvData)

	helper.Assert2DFloat32ArrayEqual(t, y, yData)
	helper.Assert2DFloat32ArrayEqual(t, u, uvExpected)
	helper.Assert2DFloat32ArrayEqual(t, v, uvExpected)
}

func TestSubsample420(t *testing.T) {
	yuvData := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	yExpected := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvExpected := data.Layer{
		{25.00, 22.50, 45.00, 32.50},
		{52.50, 40.00, 27.50, 37.50},
	}

	subsampler := Subsampler420{}
	y, u, v := subsampler.Subsample(yuvData, yuvData, yuvData)

	helper.Assert2DFloat32ArrayEqual(t, y, yExpected)
	helper.Assert2DFloat32ArrayEqual(t, u, uvExpected)
	helper.Assert2DFloat32ArrayEqual(t, v, uvExpected)
}

func TestSupersample420(t *testing.T) {
	yData := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvData := data.Layer{
		{23, 38, 55, 43},
		{46, 32, 68, 11},
	}

	uvExpected := data.Layer{
		{23, 23, 38, 38, 55, 55, 43, 43},
		{23, 23, 38, 38, 55, 55, 43, 43},
		{46, 46, 32, 32, 68, 68, 11, 11},
		{46, 46, 32, 32, 68, 68, 11, 11},
	}

	subsampler := Subsampler420{}
	y, u, v := subsampler.Supersample(yData, uvData, uvData)

	helper.Assert2DFloat32ArrayEqual(t, y, yData)
	helper.Assert2DFloat32ArrayEqual(t, u, uvExpected)
	helper.Assert2DFloat32ArrayEqual(t, v, uvExpected)
}

func TestSubsample422(t *testing.T) {
	yuvData := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	yExpected := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvExpected := data.Layer{
		{10.00, 20.00, 30.00, 45.00},
		{40.00, 25.00, 60.00, 20.00},
		{90.00, 35.00, 40.00, 40.00},
		{15.00, 45.00, 15.00, 35.00},
	}

	subsampler := Subsampler422{}
	y, u, v := subsampler.Subsample(yuvData, yuvData, yuvData)

	helper.Assert2DFloat32ArrayEqual(t, y, yExpected)
	helper.Assert2DFloat32ArrayEqual(t, u, uvExpected)
	helper.Assert2DFloat32ArrayEqual(t, v, uvExpected)
}

func TestSupersample422(t *testing.T) {
	yData := data.Layer{
		{10, 10, 20, 20, 30, 30, 40, 50},
		{30, 50, 10, 40, 80, 40, 20, 20},
		{90, 90, 10, 60, 20, 60, 10, 70},
		{10, 20, 30, 60, 10, 20, 30, 40},
	}

	uvData := data.Layer{
		{23, 38, 55, 43},
		{27, 43, 52, 66},
		{46, 32, 68, 11},
		{87, 68, 45, 33},
	}

	uvExpected := data.Layer{
		{23, 23, 38, 38, 55, 55, 43, 43},
		{27, 27, 43, 43, 52, 52, 66, 66},
		{46, 46, 32, 32, 68, 68, 11, 11},
		{87, 87, 68, 68, 45, 45, 33, 33},
	}

	subsampler := Subsampler422{}
	y, u, v := subsampler.Supersample(yData, uvData, uvData)

	helper.Assert2DFloat32ArrayEqual(t, y, yData)
	helper.Assert2DFloat32ArrayEqual(t, u, uvExpected)
	helper.Assert2DFloat32ArrayEqual(t, v, uvExpected)
}
