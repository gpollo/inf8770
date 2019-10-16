package subsampler

import (
	"errors"
	"jpeg2000/data"
)

type Subsampler interface {
	Subsample(y, u, v data.Layer) (data.Layer, data.Layer, data.Layer)
	Supersample(y, u, v data.Layer) (data.Layer, data.Layer, data.Layer)
	ToProtobuf() data.Subsampling
}

func ScaleLayers(y1, u1, v1 data.Layer, xScale, yScale int) (data.Layer, data.Layer, data.Layer) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if xScale*uSizeX != ySizeX {
		panic("Invalid X size for U layer")
	}

	if xScale*vSizeX != ySizeX {
		panic("Invalid X size for V layer")
	}

	if yScale*uSizeY != ySizeY {
		panic("Invalid Y size for U layer")
	}

	if yScale*vSizeY != ySizeY {
		panic("Invalid Y size for V layer")
	}

	y2 := y1.Copy()
	u2 := u1.ScaleInteger(xScale, yScale)
	v2 := v1.ScaleInteger(xScale, yScale)

	return y2, u2, v2
}

type Subsampler410 struct{}

func (s *Subsampler410) Subsample(y1, u1, v1 data.Layer) (data.Layer, data.Layer, data.Layer) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if ySizeX != uSizeX || ySizeX != vSizeX {
		panic("X dimensions aren't equal")
	}

	if ySizeY != uSizeY || ySizeY != vSizeY {
		panic("Y dimensions aren't equal")
	}

	sizeX := ySizeX
	sizeY := ySizeY

	y2 := data.NewLayer(sizeX, sizeY)
	u2 := data.NewLayer(sizeX/4, sizeY/2)
	v2 := data.NewLayer(sizeX/4, sizeY/2)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y2[j][i] = y1[j][i]
		}
	}

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX/4; i++ {
			u := float32(0)
			u += u1[2*j+0][4*i+0]
			u += u1[2*j+0][4*i+1]
			u += u1[2*j+0][4*i+2]
			u += u1[2*j+0][4*i+3]
			u += u1[2*j+1][4*i+0]
			u += u1[2*j+1][4*i+1]
			u += u1[2*j+1][4*i+2]
			u += u1[2*j+1][4*i+3]
			u2[j][i] = u / 8

			v := float32(0)
			v += v1[2*j+0][4*i+0]
			v += v1[2*j+0][4*i+1]
			v += v1[2*j+0][4*i+2]
			v += v1[2*j+0][4*i+3]
			v += v1[2*j+1][4*i+0]
			v += v1[2*j+1][4*i+1]
			v += v1[2*j+1][4*i+2]
			v += v1[2*j+1][4*i+3]
			v2[j][i] = v / 8
		}
	}

	return y2, u2, v2
}

func (s *Subsampler410) Supersample(y, u, v data.Layer) (data.Layer, data.Layer, data.Layer) {
	return ScaleLayers(y, u, v, 4, 2)
}

func (s *Subsampler410) ToProtobuf() data.Subsampling {
	return data.Subsampling_SUBSAMPLING_410
}

type Subsampler420 struct{}

func (s *Subsampler420) Subsample(y1, u1, v1 data.Layer) (data.Layer, data.Layer, data.Layer) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if ySizeX != uSizeX || ySizeX != vSizeX {
		panic("X dimensions aren't equal")
	}

	if ySizeY != uSizeY || ySizeY != vSizeY {
		panic("Y dimensions aren't equal")
	}

	sizeX := ySizeX
	sizeY := ySizeY

	y2 := data.NewLayer(sizeX, sizeY)
	u2 := data.NewLayer(sizeX/2, sizeY/2)
	v2 := data.NewLayer(sizeX/2, sizeY/2)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y2[j][i] = y1[j][i]
		}
	}

	for j := 0; j < sizeY/2; j++ {
		for i := 0; i < sizeX/2; i++ {
			u := float32(0)
			u += u1[2*j+0][2*i+0]
			u += u1[2*j+1][2*i+0]
			u += u1[2*j+1][2*i+1]
			u += u1[2*j+0][2*i+1]
			u2[j][i] = u / 4

			v := float32(0)
			v += v1[2*j+0][2*i+0]
			v += v1[2*j+1][2*i+0]
			v += v1[2*j+1][2*i+1]
			v += v1[2*j+0][2*i+1]
			v2[j][i] = v / 4
		}
	}

	return y2, u2, v2
}

func (s *Subsampler420) Supersample(y, u, v data.Layer) (data.Layer, data.Layer, data.Layer) {
	return ScaleLayers(y, u, v, 2, 2)
}

func (s *Subsampler420) ToProtobuf() data.Subsampling {
	return data.Subsampling_SUBSAMPLING_420
}

type Subsampler422 struct{}

func (s *Subsampler422) Subsample(y1, u1, v1 data.Layer) (data.Layer, data.Layer, data.Layer) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if ySizeX != uSizeX || ySizeX != vSizeX {
		panic("X dimensions aren't equal")
	}

	if ySizeY != uSizeY || ySizeY != vSizeY {
		panic("Y dimensions aren't equal")
	}

	sizeX := ySizeX
	sizeY := ySizeY

	y2 := data.NewLayer(sizeX, sizeY)
	u2 := data.NewLayer(sizeX/2, sizeY)
	v2 := data.NewLayer(sizeX/2, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y2[j][i] = y1[j][i]
		}
	}

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX/2; i++ {
			u := float32(0)
			u += u1[j][2*i+0]
			u += u1[j][2*i+1]
			u2[j][i] = u / 2

			v := float32(0)
			v += v1[j][2*i+0]
			v += v1[j][2*i+1]
			v2[j][i] = v / 2
		}
	}

	return y2, u2, v2
}

func (s *Subsampler422) Supersample(y, u, v data.Layer) (data.Layer, data.Layer, data.Layer) {
	return ScaleLayers(y, u, v, 2, 1)
}

func (s *Subsampler422) ToProtobuf() data.Subsampling {
	return data.Subsampling_SUBSAMPLING_422
}

type Subsampler444 struct{}

func (s *Subsampler444) Subsample(y1, u1, v1 data.Layer) (data.Layer, data.Layer, data.Layer) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if ySizeX != uSizeX || ySizeX != vSizeX {
		panic("X dimensions aren't equal")
	}

	if ySizeY != uSizeY || ySizeY != vSizeY {
		panic("Y dimensions aren't equal")
	}

	y2 := y1.Copy()
	u2 := u1.Copy()
	v2 := v1.Copy()

	return y2, u2, v2
}

func (s *Subsampler444) Supersample(y, u, v data.Layer) (data.Layer, data.Layer, data.Layer) {
	return y.Copy(), u.Copy(), v.Copy()
}

func (s *Subsampler444) ToProtobuf() data.Subsampling {
	return data.Subsampling_SUBSAMPLING_444
}

func FromProtobuf(d data.Subsampling) (Subsampler, error) {
	switch d {
	case data.Subsampling_SUBSAMPLING_410:
		return &Subsampler410{}, nil
	case data.Subsampling_SUBSAMPLING_420:
		return &Subsampler420{}, nil
	case data.Subsampling_SUBSAMPLING_422:
		return &Subsampler422{}, nil
	case data.Subsampling_SUBSAMPLING_444:
		return &Subsampler444{}, nil
	default:
		return nil, errors.New("Unknown subsampling format from protobuf")
	}
}
