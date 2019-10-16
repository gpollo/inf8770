package main

import (
	"errors"
	"fmt"
	"jpeg2000/data"
	"math"
)

type HaarWavelet struct {
	level uint32
}

func NewHaarWavelet(level int64) (*HaarWavelet, error) {
	if level < 1 {
		return nil, fmt.Errorf("Wavelet level (%d) cannot be negative or zero", level)
	}

	return &HaarWavelet{level: uint32(level)}, nil
}

func (w *HaarWavelet) GetXLowPassFilter(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			data[j][i] = (l[j][2*i+0] + l[j][2*i+1]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) GetYLowPassFilter(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][i] = (l[2*j+0][i] + l[2*j+1][i]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) GetXHighPassFilter(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			data[j][i] = (l[j][2*i+0] - l[j][2*i+1]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) GetYHighPassFilter(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX, sizeY/2)

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][i] = (l[2*j+0][i] - l[2*j+1][i]) / 2
		}
	}

	return data
}

func (w *HaarWavelet) ScaleX(f1, f2 data.Layer) data.Layer {
	sizeX1, sizeY1 := f1.GetDimensions()
	sizeX2, sizeY2 := f2.GetDimensions()

	if sizeX1 != sizeX2 || sizeY1 != sizeY2 {
		panic("Image dimensions aren't equal")
	}

	sizeX := sizeX1
	sizeY := sizeY1
	data := data.NewLayer(2*sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			data[j][2*i+0] = f1[j][i] + f2[j][i]
			data[j][2*i+1] = f1[j][i] - f2[j][i]
		}
	}

	return data
}

func (w *HaarWavelet) ScaleY(f1, f2 data.Layer) data.Layer {
	sizeX1, sizeY1 := f1.GetDimensions()
	sizeX2, sizeY2 := f2.GetDimensions()

	if sizeX1 != sizeX2 || sizeY1 != sizeY2 {
		panic("Image dimensions aren't equal")
	}

	sizeX := sizeX1
	sizeY := sizeY1
	data := data.NewLayer(sizeX, 2*sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			data[2*j+0][i] = f1[j][i] + f2[j][i]
			data[2*j+1][i] = f1[j][i] - f2[j][i]
		}
	}

	return data
}

func (w *HaarWavelet) CopyIntoQuadrant(from, into data.Layer, quadrant int) {
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

func (w *HaarWavelet) CopyFromQuadrant(from, into data.Layer, quadrant int) {
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

func (w *HaarWavelet) WaveletTransform(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX, sizeY)

	level := int(w.level)
	if level == 0 {
		level = 2
	}

	for i := 0; i < level; i++ {
		fll := w.GetYLowPassFilter(w.GetXLowPassFilter(l))
		flh := w.GetYHighPassFilter(w.GetXLowPassFilter(l))
		fhl := w.GetYLowPassFilter(w.GetXHighPassFilter(l))
		fhh := w.GetYHighPassFilter(w.GetXHighPassFilter(l))

		w.CopyIntoQuadrant(flh, data, 1)
		w.CopyIntoQuadrant(fll, data, 2)
		w.CopyIntoQuadrant(fhl, data, 3)
		w.CopyIntoQuadrant(fhh, data, 4)

		l = fll
	}

	return data
}

func (w *HaarWavelet) WaveletInverse(l data.Layer) data.Layer {
	copy := l.Copy()
	sizeX, sizeY := copy.GetDimensions()

	level := int(w.level)
	if level == 0 {
		level = 2
	}

	for i := level; i > 0; i-- {
		factor := int(math.Pow(2.0, float64(i)))
		fll := data.NewLayer(sizeX/factor, sizeY/factor)
		flh := data.NewLayer(sizeX/factor, sizeY/factor)
		fhl := data.NewLayer(sizeX/factor, sizeY/factor)
		fhh := data.NewLayer(sizeX/factor, sizeY/factor)

		w.CopyFromQuadrant(copy, flh, 1)
		w.CopyFromQuadrant(copy, fll, 2)
		w.CopyFromQuadrant(copy, fhl, 3)
		w.CopyFromQuadrant(copy, fhh, 4)

		fl := w.ScaleY(fll, flh)
		fh := w.ScaleY(fhl, fhh)
		f := w.ScaleX(fl, fh)

		if i == 1 {
			copy = f
		} else {
			w.CopyIntoQuadrant(f, copy, 2)
		}
	}

	return copy
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

func (w *HaarWavelet) FromProtobuf(d *data.WaveletConfig) error {
	c := d.GetHaar()
	if c == nil {
		return errors.New("Could not deserialize haar wavelet from protobuf data")
	}

	w.level = c.Level

	return nil
}
