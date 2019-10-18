package data

import (
	"errors"
	"image"
	"image/color"
	"math"
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
		return Image{}, errors.New("Mismatch in size X between layers")
	}

	if rSizeY != gSizeY || rSizeY != bSizeY {
		return Image{}, errors.New("Mismatch in size Y between layers")
	}

	sizeX := rSizeX
	sizeY := rSizeY

	y := NewLayer(sizeX, sizeY)
	u := NewLayer(sizeX, sizeY)
	v := NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y[j][i] = (r[j][i] + 2*g[j][i] + b[j][i]) / 4
			u[j][i] = (b[j][i] - g[j][i])
			v[j][i] = (r[j][i] - g[j][i])
		}
	}

	return Image{y, u, v}, nil
}

func (i Image) YUVToRGB() (Image, error) {
	y := i[0]
	u := i[1]
	v := i[2]

	ySizeX, ySizeY := y.GetDimensions()
	uSizeX, uSizeY := u.GetDimensions()
	vSizeX, vSizeY := v.GetDimensions()

	if ySizeX != uSizeX || ySizeX != vSizeX {
		return Image{}, errors.New("Mismatch in size X between layers")
	}

	if ySizeY != uSizeY || ySizeY != vSizeY {
		return Image{}, errors.New("Mismatch in size Y between layers")
	}

	sizeX := ySizeX
	sizeY := ySizeY

	r := NewLayer(sizeX, sizeY)
	g := NewLayer(sizeX, sizeY)
	b := NewLayer(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			g[j][i] = y[j][i] - ((u[j][i] + v[j][i]) / 4)
			r[j][i] = v[j][i] + g[j][i]
			b[j][i] = u[j][i] + g[j][i]

			g[j][i] = float32(math.Min(float64(g[j][i]), float64(0xffff)))
			r[j][i] = float32(math.Min(float64(r[j][i]), float64(0xffff)))
			b[j][i] = float32(math.Min(float64(b[j][i]), float64(0xffff)))
		}
	}

	return Image{r, g, b}, nil
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
	r := float64(i[0][y][x])
	g := float64(i[1][y][x])
	b := float64(i[2][y][x])

	r = math.Min(float64(0xffff), r)
	g = math.Min(float64(0xffff), g)
	b = math.Min(float64(0xffff), b)

	r = math.Max(float64(0x0000), r)
	g = math.Max(float64(0x0000), g)
	b = math.Max(float64(0x0000), b)

	return &color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: 0xffff,
	}
}
