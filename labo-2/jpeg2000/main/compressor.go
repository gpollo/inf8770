package main

import "jpeg2000/data"

type Compressor interface {
	CompressLayer(d ImageData) data.FileImageLayer
	DecompressLayer(d data.FileImageLayer) ImageData
}
