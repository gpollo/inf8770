package main

import (
	"fmt"
	"image"
)

type ImageData [][]byte

func NewImageData(sizeX, sizeY int) ImageData {
	data := make([][]byte, sizeY)
	for j := 0; j < sizeY; j++ {
		data[j] = make([]byte, sizeX)
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
			rData[j][i] = byte(r)
			gData[j][i] = byte(g)
			bData[j][i] = byte(b)
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
