package wavelet

import (
	"errors"
	"jpeg2000/data"
)

type DummyWavelet struct {
}

func NewDummyWavelet() *DummyWavelet {
	return &DummyWavelet{}
}

func (w *DummyWavelet) WaveletTransform(l data.Layer) data.Layer {
	return l.Copy()
}

func (w *DummyWavelet) WaveletInverse(l data.Layer) data.Layer {
	return l.Copy()
}

func (w *DummyWavelet) ToProtobuf() *data.WaveletConfig {
	return &data.WaveletConfig{
		Data: &data.WaveletConfig_Dummy{
			Dummy: &data.WaveletDummy{},
		},
	}
}

func (w *DummyWavelet) FromProtobuf(d *data.WaveletConfig) error {
	c := d.GetDummy()
	if c == nil {
		return errors.New("Could not deserialize dummy wavelet from protobuf data")
	}

	return nil
}
