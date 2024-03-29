package wavelet

import (
	"errors"
	"fmt"
	"jpeg2000/data"
)

type DaubechiesWavelet struct {
	level       uint32
	coefficient uint32
}

func NewDaubechiesWavelet(level int64, coefficient int64) (*DaubechiesWavelet, error) {
	if level < 1 {
		return nil, fmt.Errorf("Wavelet level (%d) cannot be negative or zero", level)
	}

	if coefficient < 1 {
		return nil, fmt.Errorf("Daubechies coefficient (%d) cannot be negative or zero", coefficient)
	}

	return &DaubechiesWavelet{level: uint32(level), coefficient: uint32(coefficient)}, nil
}

func (w *DaubechiesWavelet) SetLevel(level uint32) {
	if level < 2 {
		return
	}

	w.level = level
}

func (w *DaubechiesWavelet) WaveletTransform(l data.Layer) data.Layer {
	pyWavelet := PyWavelet{mode: fmt.Sprintf("db%d", w.coefficient)}
	return pyWavelet.WaveletTransform(l)
}

func (w *DaubechiesWavelet) WaveletInverse(l data.Layer) data.Layer {
	pyWavelet := PyWavelet{mode: fmt.Sprintf("db%d", w.coefficient)}
	return pyWavelet.WaveletInverse(l)
}

func (w *DaubechiesWavelet) ToProtobuf() *data.WaveletConfig {
	return &data.WaveletConfig{
		Data: &data.WaveletConfig_Daubechies{
			Daubechies: &data.WaveletDaubechies{
				Level:       w.level,
				Coefficient: w.coefficient,
			},
		},
	}
}

func (w *DaubechiesWavelet) FromProtobuf(d *data.WaveletConfig) error {
	c := d.GetDaubechies()
	if c == nil {
		return errors.New("Could not deserialize daubechies wavelet from protobuf data")
	}

	w.level = c.Level
	w.coefficient = c.Coefficient

	return nil
}
