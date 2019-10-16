package data

import (
	"errors"
	"image"
	"image/color"
)

type Image [3]Layer

func (i Image) RGBToYUV() (Image, error) {
	r := i[0]
	g := i[1]
	b := i[2]

	rSizeX, rSizeY := r.GetDimensions()
	gSizeX, gSizeY := g.GetDimensions()
	bSizeX, bSizeY := b.GetDimensions()

	if rSizeX != gSizeX || rSizeX != bSizeX {
		return Image{}, errors.New("X dimensions aren't equal")
	}

	if rSizeY != gSizeY || rSizeY != bSizeY {
		return Image{}, errors.New("Y dimensions aren't equal")
	}

	sizeX := rSizeX
	sizeY := rSizeY

	y := NewLayer(sizeX, sizeY)
	u := NewLayer(sizeX, sizeY)
	v := NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y[j][i] = (r[j][i] + 2*g[j][i] + b[j][i]) / 4
			u[j][i] = b[j][i] + g[j][i]
			v[j][i] = r[j][i] + g[j][i]
		}
	}

	return Image{y, u, v}, nil
}

func (i Image) ConvertYUVToRGB() (Image, error) {
	y := i[0]
	u := i[1]
	v := i[2]

	ySizeX, ySizeY := y.GetDimensions()
	uSizeX, uSizeY := u.GetDimensions()
	vSizeX, vSizeY := v.GetDimensions()

	if ySizeX != uSizeX || ySizeX != vSizeX {
		return Image{}, errors.New("X dimensions aren't equal")
	}

	if ySizeY != uSizeY || ySizeY != vSizeY {
		return Image{}, errors.New("Y dimensions aren't equal")
	}

	sizeX := ySizeX
	sizeY := ySizeY

	r := NewLayer(sizeX, sizeY)
	g := NewLayer(sizeX, sizeY)
	b := NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			g[j][i] = y[j][i] - (u[j][i]+v[j][i])/4
			r[j][i] = v[j][i] + g[j][i]
			b[j][i] = u[j][i] + g[j][i]
		}
	}

	return Image{r, b, g}, nil
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	sizeX, sizeY := i[0].GetDimensions()

	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: sizeX,
			Y: sizeY,
		},
	}
}

func (i *Image) At(x, y int) color.Color {
	return &color.RGBA64{
		R: uint16(i[0][y][x]),
		G: uint16(i[1][y][x]),
		B: uint16(i[2][y][x]),
		A: 0xffff,
	}
}
