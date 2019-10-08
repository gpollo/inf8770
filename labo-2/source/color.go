package main

func ConvertRGBToYUV(r, g, b ImageData) (ImageData, ImageData, ImageData) {
	rSizeX, rSizeY := r.GetDimensions()
	gSizeX, gSizeY := g.GetDimensions()
	bSizeX, bSizeY := b.GetDimensions()

	if rSizeX != gSizeX || rSizeX != bSizeX {
		panic("X dimensions aren't equal")
	}

	if rSizeY != gSizeY || rSizeY != bSizeY {
		panic("Y dimensions aren't equal")
	}

	sizeX := rSizeX
	sizeY := rSizeY

	y := NewImageData(sizeX, sizeY)
	u := NewImageData(sizeX, sizeY)
	v := NewImageData(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y[j][i] = (r[j][i] + 2*g[j][i] + b[j][i]) / 4
			u[j][i] = b[j][i] + g[j][i]
			v[j][i] = r[j][i] + g[j][i]
		}
	}

	return y, u, v
}

func ConvertYUVToRGB(y, u, v ImageData) (ImageData, ImageData, ImageData) {
	ySizeX, ySizeY := y.GetDimensions()
	uSizeX, uSizeY := u.GetDimensions()
	vSizeX, vSizeY := v.GetDimensions()

	if ySizeX != uSizeX || ySizeX != vSizeX {
		panic("X dimensions aren't equal")
	}

	if ySizeY != uSizeY || ySizeY != vSizeY {
		panic("Y dimensions aren't equal")
	}

	sizeX := ySizeX
	sizeY := ySizeY

	r := NewImageData(sizeX, sizeY)
	g := NewImageData(sizeX, sizeY)
	b := NewImageData(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			g[j][i] = y[j][i] - (u[j][i]+v[j][i])/4
			r[j][i] = v[j][i] + g[j][i]
			b[j][i] = u[j][i] + g[j][i]
		}
	}

	return r, b, g
}
