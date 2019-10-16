package main

import (
	"errors"
	"fmt"
	"jpeg2000/data"
	"math"
)

type DeadZoneQuantifier struct {
	width  uint32
	delta  uint32
	offset float32
}

func NewDeadZoneQuantifier(width, delta int64, offset float64) (*DeadZoneQuantifier, error) {
	if width < 0 {
		return nil, fmt.Errorf("Dead zone width (%d) cannot be negative", width)
	}

	if delta <= 0 {
		return nil, fmt.Errorf("Dead zone delta (%d) cannot be negative or zero", delta)
	}

	return &DeadZoneQuantifier{
		width:  uint32(width),
		delta:  uint32(delta),
		offset: float32(offset),
	}, nil
}

func (q *DeadZoneQuantifier) QuantifierTransform(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			var sign float64
			if l[j][i] < 0 {
				sign = -1.0
			} else {
				sign = 1.0
			}

			value := math.Abs(float64(l[j][i]))
			width := float64(q.width)
			delta := float64(q.delta)

			data[j][i] = float32(sign * math.Max(0, math.Floor(((value-width)/delta)+1)))
		}
	}

	return data
}

func (q *DeadZoneQuantifier) QuantifierInverse(l data.Layer) data.Layer {
	sizeX, sizeY := l.GetDimensions()
	data := data.NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			var sign float32
			if l[j][i] < 0 {
				sign = -1.0
			} else {
				sign = 1.0
			}

			value := float32(math.Abs(float64(l[j][i])))
			width := float32(q.width)
			delta := float32(q.delta)

			data[j][i] = sign * (width + delta*(value-1+q.offset))
		}
	}

	return data
}

func (q *DeadZoneQuantifier) ToProtobuf() *data.QuantifierConfig {
	return &data.QuantifierConfig{
		Data: &data.QuantifierConfig_DeadZone{
			DeadZone: &data.QuantifierDeadZone{
				Width:  q.width,
				Delta:  q.delta,
				Offset: q.offset,
			},
		},
	}
}

func (q *DeadZoneQuantifier) FromProtobuf(d *data.QuantifierConfig) error {
	c := d.GetDeadZone()
	if c == nil {
		return errors.New("Could not deserialize dead zone quantifier from protobuf data")
	}

	q.width = c.Width
	q.delta = c.Delta
	q.offset = c.Offset

	return nil
}
