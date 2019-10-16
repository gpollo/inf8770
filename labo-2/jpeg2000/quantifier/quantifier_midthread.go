package quantifier

import (
	"errors"
	"fmt"
	"jpeg2000/data"
	"math"
)

type MidThreadQuantifier struct {
	delta uint32
}

func NewMidThreadQuantifier(delta int64) (*MidThreadQuantifier, error) {
	if delta <= 0 {
		return nil, fmt.Errorf("Mid-thread delta (%d) cannot be negative or zero", delta)
	}

	return &MidThreadQuantifier{delta: uint32(delta)}, nil
}

func (q *MidThreadQuantifier) QuantifierTransform(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			value := math.Abs(float64(l[j][i]))
			delta := float64(q.delta)

			data[j][i] = float32(math.Round(value / delta))
		}
	}

	return data
}

func (q *MidThreadQuantifier) QuantifierInverse(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			value := float32(l[j][i])
			delta := float32(q.delta)

			data[j][i] = delta * (value + 0.5)
		}
	}

	return data
}

func (q *MidThreadQuantifier) ToProtobuf() *data.QuantifierConfig {
	return &data.QuantifierConfig{
		Data: &data.QuantifierConfig_MidThread{
			MidThread: &data.QuantifierMidThread{
				Delta: q.delta,
			},
		},
	}
}

func (q *MidThreadQuantifier) FromProtobuf(d *data.QuantifierConfig) error {
	c := d.GetMidThread()
	if c == nil {
		return errors.New("Could not deserialize mid-thread quantifier from protobuf data")
	}

	q.delta = c.Delta

	return nil
}
