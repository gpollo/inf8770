package main

import "source/data"

type Compressor interface {
	CompressLayer(d ImageData) data.FileImageLayer
	DecompressLayer(d data.FileImageLayer) ImageData
}
