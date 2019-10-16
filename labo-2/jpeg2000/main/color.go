package main

import "jpeg2000/data"

func ConvertRGBToYUV(r, g, b data.Layer) (data.Layer, data.Layer, data.Layer) {
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

	y := data.NewLayer(sizeX, sizeY)
	u := data.NewLayer(sizeX, sizeY)
	v := data.NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y[j][i] = (r[j][i] + 2*g[j][i] + b[j][i]) / 4
			u[j][i] = b[j][i] + g[j][i]
			v[j][i] = r[j][i] + g[j][i]
		}
	}

	return y, u, v
}

func ConvertYUVToRGB(y, u, v data.Layer) (data.Layer, data.Layer, data.Layer) {
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

	r := data.NewLayer(sizeX, sizeY)
	g := data.NewLayer(sizeX, sizeY)
	b := data.NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			g[j][i] = y[j][i] - (u[j][i]+v[j][i])/4
			r[j][i] = v[j][i] + g[j][i]
			b[j][i] = u[j][i] + g[j][i]
		}
	}

	return r, b, g
}
