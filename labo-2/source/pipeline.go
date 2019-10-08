package main

import (
	"image"
)

func EncodeImage(input image.Image) ([]byte, error) {
	GetImageData(input)
	return []byte{}, nil
}

func DecodeImage(input []byte) (image.Image, error) {
	return nil, nil
}
