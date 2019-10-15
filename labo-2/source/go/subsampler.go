package main

type subsampler interface {
	Subsample(y, u, v ImageData) (ImageData, ImageData, ImageData)
	Supersample(y, u, v ImageData) (ImageData, ImageData, ImageData)
}

type Subsampler410 struct{}

func (s *Subsampler410) Subsample(y1, u1, v1 ImageData) (ImageData, ImageData, ImageData) {
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

	y2 := NewImageData(sizeX, sizeY)
	u2 := NewImageData(sizeX/4, sizeY/2)
	v2 := NewImageData(sizeX/4, sizeY/2)

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

func (s *Subsampler410) Supersample(y1, u1, v1 ImageData) (ImageData, ImageData, ImageData) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if 4*uSizeX != ySizeX {
		panic("Invalid X size for U layer")
	}

	if 4*vSizeX != ySizeX {
		panic("Invalid X size for V layer")
	}

	if 2*uSizeY != ySizeY {
		panic("Invalid Y size for U layer")
	}

	if 2*vSizeY != ySizeY {
		panic("Invalid Y size for V layer")
	}

	y2 := y1.Copy()
	u2 := u1.ScaleInteger(4, 2)
	v2 := v1.ScaleInteger(4, 2)

	return y2, u2, v2
}

type Subsampler420 struct{}

func (s *Subsampler420) Subsample(y1, u1, v1 ImageData) (ImageData, ImageData, ImageData) {
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

	y2 := NewImageData(sizeX, sizeY)
	u2 := NewImageData(sizeX/2, sizeY/2)
	v2 := NewImageData(sizeX/2, sizeY/2)

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

func (s *Subsampler420) Supersample(y1, u1, v1 ImageData) (ImageData, ImageData, ImageData) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if 2*uSizeX != ySizeX {
		panic("Invalid X size for U layer")
	}

	if 2*vSizeX != ySizeX {
		panic("Invalid X size for V layer")
	}

	if 2*uSizeY != ySizeY {
		panic("Invalid Y size for U layer")
	}

	if 2*vSizeY != ySizeY {
		panic("Invalid Y size for V layer")
	}

	y2 := y1.Copy()
	u2 := u1.ScaleInteger(2, 2)
	v2 := v1.ScaleInteger(2, 2)

	return y2, u2, v2
}

type Subsampler422 struct{}

func (s *Subsampler422) Subsample(y1, u1, v1 ImageData) (ImageData, ImageData, ImageData) {
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

	y2 := NewImageData(sizeX, sizeY)
	u2 := NewImageData(sizeX/2, sizeY)
	v2 := NewImageData(sizeX/2, sizeY)

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

func (s *Subsampler422) Supersample(y1, u1, v1 ImageData) (ImageData, ImageData, ImageData) {
	ySizeX, ySizeY := y1.GetDimensions()
	uSizeX, uSizeY := u1.GetDimensions()
	vSizeX, vSizeY := v1.GetDimensions()

	if 2*uSizeX != ySizeX {
		panic("Invalid X size for U layer")
	}

	if 2*vSizeX != ySizeX {
		panic("Invalid X size for V layer")
	}

	if uSizeY != ySizeY {
		panic("Invalid Y size for U layer")
	}

	if vSizeY != ySizeY {
		panic("Invalid Y size for V layer")
	}

	y2 := y1.Copy()
	u2 := u1.ScaleInteger(2, 1)
	v2 := v1.ScaleInteger(2, 1)

	return y2, u2, v2
}
