package main

import (
	"fmt"
	"jpeg2000/data"
)

type DaubechiesWavelet struct {
	level uint32
}

func NewDaubechiesWavelet() DaubechiesWavelet {
	return DaubechiesWavelet{level: 2}
}

func (w *DaubechiesWavelet) SetLevel(level uint32) {
	if level < 2 {
		return
	}

	w.level = level
}

func (w *DaubechiesWavelet) WaveletTransform(d ImageData) ImageData {
	pyWavelet := PyWavelet{mode: fmt.Sprintf("db%d", w.level)}
	return pyWavelet.WaveletTransform(d)
}

func (w *DaubechiesWavelet) WaveletInverse(d ImageData) ImageData {
	pyWavelet := PyWavelet{mode: fmt.Sprintf("db%d", w.level)}
	return pyWavelet.WaveletInverse(d)
}

func (w *DaubechiesWavelet) ToProtobuf() *data.WaveletConfig {
	return &data.WaveletConfig{
		Data: &data.WaveletConfig_Daubechies{
			Daubechies: &data.WaveletDaubechies{
				Level: w.level,
			},
		},
	}
}

func (w *DaubechiesWavelet) FromProtobuf(d data.WaveletConfig) {
	c := d.GetDaubechies()
	if c == nil {
		panic("Could not deserialize daubechies wavelet from protobuf")
	}

	w.level = c.Level
}
