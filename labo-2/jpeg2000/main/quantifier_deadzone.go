package main

import (
	"jpeg2000/data"
	"math"
)

type DeadZoneQuantifier struct {
	width  uint32
	delta  uint32
	offset float32
}

func (q *DeadZoneQuantifier) QuantifierTransform(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			var sign float64
			if d[j][i] < 0 {
				sign = -1.0
			} else {
				sign = 1.0
			}

			value := math.Abs(float64(d[j][i]))
			width := float64(q.width)
			delta := float64(q.delta)

			data[j][i] = float32(sign * math.Max(0, math.Floor(((value-width)/delta)+1)))
		}
	}

	return data
}

func (q *DeadZoneQuantifier) QuantifierInverse(d ImageData) ImageData {
	sizeX, sizeY := d.GetDimensions()
	data := NewImageData(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			var sign float32
			if d[j][i] < 0 {
				sign = -1.0
			} else {
				sign = 1.0
			}

			value := float32(math.Abs(float64(d[j][i])))
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

func (q *DeadZoneQuantifier) FromProtobuf(d data.QuantifierConfig) {
	c := d.GetDeadZone()
	if c == nil {
		panic("Could not deserialize dead zone quantifier from protobuf")
	}

	q.width = c.Width
	q.delta = c.Delta
	q.offset = c.Offset
}
