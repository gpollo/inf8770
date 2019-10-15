package main

import (
	"jpeg2000/data"
	"math"
)

type HaarWavelet struct {
	level uint32
}

func (w *HaarWavelet) GetXLowPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			data[j][i] = (d[j][2*i+0] + d[j][2*i+1]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) GetYLowPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][i] = (d[2*j+0][i] + d[2*j+1][i]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) GetXHighPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			data[j][i] = (d[j][2*i+0] - d[j][2*i+1]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) GetYHighPassFilter(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][i] = (d[2*j+0][i] - d[2*j+1][i]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) ScaleX(f1, f2 ImageData) ImageData {
	sizeX1, sizeY1 := f1.GetDimensions()
	sizeX2, sizeY2 := f2.GetDimensions()

	if sizeX1 != sizeX2 || sizeY1 != sizeY2 {
		panic("Image dimensions aren't equal")
	}

	sizeX := sizeX1
	sizeY := sizeY1
	data := NewImageData(2*sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][2*i+0] = f1[j][i] + f2[j][i]
			data[j][2*i+1] = f1[j][i] - f2[j][i]
		}
	}

	return data
}

func (w *HaarWavelet) ScaleY(f1, f2 ImageData) ImageData {
	sizeX1, sizeY1 := f1.GetDimensions()
	sizeX2, sizeY2 := f2.GetDimensions()

	if sizeX1 != sizeX2 || sizeY1 != sizeY2 {
		panic("Image dimensions aren't equal")
	}

	sizeX := sizeX1
	sizeY := sizeY1
	data := NewImageData(sizeX, 2*sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			data[2*j+0][i] = f1[j][i] + f2[j][i]
			data[2*j+1][i] = f1[j][i] - f2[j][i]
		}
	}

	return data
}

func (w *HaarWavelet) CopyIntoQuadrant(from, into ImageData, quadrant int) {
	sizeFromX, sizeFromY := from.GetDimensions()
	sizeIntoX, sizeIntoY := into.GetDimensions()

	if 2*sizeFromX > sizeIntoX {
		panic("Invalid X size for copying from")
	}

	if 2*sizeFromY > sizeIntoY {
		panic("Invalid Y size for copying from")
	}

	var offsetX int
	var offsetY int
	switch quadrant {
	case 1:
		offsetX = sizeFromX
		offsetY = 0
	case 2:
		offsetX = 0
		offsetY = 0
	case 3:
		offsetX = 0
		offsetY = sizeFromY
	case 4:
		offsetX = sizeFromX
		offsetY = sizeFromY
	default:
		panic("Invalid quadrant selected for copying")
	}

	for j := 0; j < sizeFromY; j++ {
		for i := 0; i < sizeFromX; i++ {
			into[j+offsetY][i+offsetX] = from[j][i]
		}
	}
}

func (w *HaarWavelet) CopyFromQuadrant(from, into ImageData, quadrant int) {
	sizeFromX, sizeFromY := from.GetDimensions()
	sizeIntoX, sizeIntoY := into.GetDimensions()

	if 2*sizeIntoX > sizeFromX {
		panic("Invalid X size for copying into")
	}

	if 2*sizeIntoY > sizeFromY {
		panic("Invalid Y size for copying into")
	}

	var offsetX int
	var offsetY int
	switch quadrant {
	case 1:
		offsetX = sizeIntoX
		offsetY = 0
	case 2:
		offsetX = 0
		offsetY = 0
	case 3:
		offsetX = 0
		offsetY = sizeIntoY
	case 4:
		offsetX = sizeIntoX
		offsetY = sizeIntoY
	default:
		panic("Invalid quadrant selected for copying")
	}

	for j := 0; j < sizeIntoY; j++ {
		for i := 0; i < sizeIntoX; i++ {
			into[j][i] = from[j+offsetY][i+offsetX]
		}
	}
}

func (w *HaarWavelet) WaveletTransform(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY)

	level := int(w.level)
	if level == 0 {
		level = 2
	}

	for i := 0; i < level; i++ {
		fll := w.GetYLowPassFilter(w.GetXLowPassFilter(d))
		flh := w.GetYHighPassFilter(w.GetXLowPassFilter(d))
		fhl := w.GetYLowPassFilter(w.GetXHighPassFilter(d))
		fhh := w.GetYHighPassFilter(w.GetXHighPassFilter(d))

		w.CopyIntoQuadrant(flh, data, 1)
		w.CopyIntoQuadrant(fll, data, 2)
		w.CopyIntoQuadrant(fhl, data, 3)
		w.CopyIntoQuadrant(fhh, data, 4)

		d = fll
	}

	return data
}

func (w *HaarWavelet) WaveletInverse(d ImageData) ImageData {
	data := d.Copy()
	sizeX, sizeY := data.GetDimensions()

	level := int(w.level)
	if level == 0 {
		level = 2
	}

	for i := level; i > 0; i-- {
		factor := int(math.Pow(2.0, float64(i)))
		fll := NewImageData(sizeX/factor, sizeY/factor)
		flh := NewImageData(sizeX/factor, sizeY/factor)
		fhl := NewImageData(sizeX/factor, sizeY/factor)
		fhh := NewImageData(sizeX/factor, sizeY/factor)

		w.CopyFromQuadrant(data, flh, 1)
		w.CopyFromQuadrant(data, fll, 2)
		w.CopyFromQuadrant(data, fhl, 3)
		w.CopyFromQuadrant(data, fhh, 4)

		fl := w.ScaleY(fll, flh)
		fh := w.ScaleY(fhl, fhh)
		f := w.ScaleX(fl, fh)

		if i == 1 {
			data = f
		} else {
			w.CopyIntoQuadrant(f, data, 2)
		}
	}

	return data
}

func (w *HaarWavelet) ToProtobuf() *data.WaveletConfig {
	return &data.WaveletConfig{
		Data: &data.WaveletConfig_Haar{
			Haar: &data.WaveletHaar{
				Level: w.level,
			},
		},
	}
}

func (w *HaarWavelet) FromProtobuf(d data.WaveletConfig) {
	c := d.GetHaar()
	if c == nil {
		panic("Could not deserialize haar wavelet from protobuf")
	}

	w.level = c.Level
}
