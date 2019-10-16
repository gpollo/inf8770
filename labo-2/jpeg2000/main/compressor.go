package main

import "jpeg2000/data"

type Compressor interface {
	CompressLayer(d data.Layer) *data.FileImageLayer
	DecompressLayer(d *data.FileImageLayer) data.Layer
}
