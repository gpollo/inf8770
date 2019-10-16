package main

import (
	"fmt"
	"image"
	"jpeg2000/data"
)

type ImageData [][]float32

func NewImageData(sizeX, sizeY int) ImageData {
	data := make([][]float32, sizeY)
	for j := 0; j < sizeY; j++ {
		data[j] = make([]float32, sizeX)
	}

	return data
}

func GetImageData(image image.Image) (ImageData, ImageData, ImageData, uint, uint) {
	sizeX := image.Bounds().Max.X - image.Bounds().Min.X
	sizeY := image.Bounds().Max.Y - image.Bounds().Min.Y

	rData := NewImageData(sizeX, sizeY)
	gData := NewImageData(sizeX, sizeY)
	bData := NewImageData(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			r, g, b, _ := image.At(i, j).RGBA()
			rData[j][i] = float32(r)
			gData[j][i] = float32(g)
			bData[j][i] = float32(b)
		}
	}

	return rData, gData, bData, uint(sizeX), uint(sizeY)
}

func RGBToYUV(r, g, b ImageData) (ImageData, ImageData, ImageData) {
	sizeXr, sizeYr := r.GetDimensions()
	sizeXg, sizeYg := g.GetDimensions()
	sizeXb, sizeYb := b.GetDimensions()

	if sizeXr != sizeXg || sizeXr != sizeXb {
		panic("Mismatch in size between layers")
	}

	if sizeYr != sizeYg || sizeYr != sizeYb {
		panic("Mismatch in size between layers")
	}

	sizeX := sizeXr
	sizeY := sizeYr

	y := NewImageData(sizeX, sizeY)
	u := NewImageData(sizeX, sizeY)
	v := NewImageData(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			y[j][i] = (r[j][i] + 2*g[j][i] + b[j][i]) / 4
			u[j][i] = (b[j][i] - g[j][i])
			v[j][i] = (r[j][i] - g[j][i])
		}
	}

	return y, u, v
}

func YUVToRGB(y, u, v ImageData) (ImageData, ImageData, ImageData) {
	sizeXy, sizeYy := y.GetDimensions()
	sizeXu, sizeYu := u.GetDimensions()
	sizeXv, sizeYv := v.GetDimensions()

	if sizeXy != sizeXu || sizeXy != sizeXv {
		panic("Mismatch in size between layers")
	}

	if sizeYy != sizeYu || sizeYy != sizeYv {
		panic("Mismatch in size between layers")
	}

	sizeX := sizeXy
	sizeY := sizeYy

	r := NewImageData(sizeX, sizeY)
	g := NewImageData(sizeX, sizeY)
	b := NewImageData(sizeX, sizeY)

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			g[j][i] = y[j][i] - ((u[j][i] + v[j][i]) / 4)
			r[j][i] = v[j][i] + g[j][i]
			b[j][i] = u[j][i] + g[j][i]

			if g[j][i] > float32(0xffff) {
				g[j][i] = float32(0xffff)
			}

			if r[j][i] > float32(0xffff) {
				r[j][i] = float32(0xffff)
			}

			if b[j][i] > float32(0xffff) {
				b[j][i] = float32(0xffff)
			}
		}
	}

	return r, g, b
}

func (d ImageData) GetDimensions() (int, int) {
	if len(d) == 0 {
		panic("Invalid image data")
	}

	return len(d[0]), len(d)
}

func (d ImageData) Copy() ImageData {
	sizeX, sizeY := d.GetDimensions()

	copy := NewImageData(sizeX, sizeY)
	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			copy[j][i] = d[j][i]
		}
	}

	return copy
}

func (d ImageData) Times(factor float32) {
	sizeX, sizeY := d.GetDimensions()

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			d[j][i] = factor * d[j][i]
		}
	}
}

func (d ImageData) ScaleInteger(scaleX, scaleY int) ImageData {
	sizeX, sizeY := d.GetDimensions()

	scaled := NewImageData(sizeX*scaleX, sizeY*scaleY)
	for j := 0; j < sizeY*scaleY; j++ {
		for i := 0; i < sizeX*scaleX; i++ {
			scaled[j][i] = d[j/scaleY][i/scaleX]
		}
	}

	return scaled
}

func (d ImageData) ToProtobuf() data.ImageData {
	sizeX, sizeY := d.GetDimensions()

	rows := make([]*data.ImageRow, sizeY)
	for j := 0; j < sizeY; j++ {
		row := data.ImageRow{Values: make([]float32, sizeX)}
		for i := 0; i < sizeX; i++ {
			row.Values[i] = d[j][i]
		}
		rows[j] = &row
	}

	return data.ImageData{Rows: rows}
}

func (d *ImageData) FromProtobuf(data data.ImageData) {
	sizeX := len(data.Rows[0].Values)
	sizeY := len(data.Rows)

	(*d) = NewImageData(sizeX, sizeY)
	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			(*d)[j][i] = data.Rows[j].Values[i]
		}
	}
}
