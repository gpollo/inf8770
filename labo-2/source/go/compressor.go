package main

type Compressor interface {
	CompressLayer(d ImageData) FileImageLayer
	DecompressLayer(d FileImageLayer) ImageData
}
