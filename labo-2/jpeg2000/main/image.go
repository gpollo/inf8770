package main

import (
	"image"
	"image/color"
	"jpeg2000/data"
)

type Image [3]data.Layer

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
