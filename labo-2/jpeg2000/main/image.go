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

func GetImageData(image image.Image) (ImageData, ImageData, ImageData) {
	sizeX := image.Bounds().Max.X - image.Bounds().Min.X
	sizeY := image.Bounds().Max.Y - image.Bounds().Min.Y

	fmt.Printf("Image X length: %d\n", sizeX)
	fmt.Printf("Image Y length: %d\n", sizeY)

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

	return rData, gData, bData
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
