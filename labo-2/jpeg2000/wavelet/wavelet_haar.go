package wavelet

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

		copyIntoQuadrant(flh, data, 1)
		copyIntoQuadrant(fll, data, 2)
		copyIntoQuadrant(fhl, data, 3)
		copyIntoQuadrant(fhh, data, 4)

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

		copyFromQuadrant(copy, flh, 1)
		copyFromQuadrant(copy, fll, 2)
		copyFromQuadrant(copy, fhl, 3)
		copyFromQuadrant(copy, fhh, 4)

		fl := scaleY(fll, flh)
		fh := scaleY(fhl, fhh)
		f := scaleX(fl, fh)

		if i == 1 {
			copy = f
		} else {
			copyIntoQuadrant(f, copy, 2)
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
