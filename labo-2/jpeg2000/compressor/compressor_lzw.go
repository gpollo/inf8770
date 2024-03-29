package compressor

import (
	"bufio"
	"bytes"
	"compress/lzw"
	"encoding/binary"
	"io"
	"jpeg2000/data"
	"jpeg2000/helper"
)

type LZWCompressor struct{}

func NewLZWCompressor() *LZWCompressor {
	return &LZWCompressor{}
}

func (c *LZWCompressor) CompressLayer(d data.Layer) *data.FileImageLayer {
	sizeX, sizeY := d.GetDimensions()

	encodedImage := data.FileImageLayer{Rows: make([][]byte, sizeY)}
	for j := 0; j < sizeY; j++ {
		encodedBuffer := bytes.NewBuffer([]byte{})
		compressor := lzw.NewWriter(encodedBuffer, lzw.MSB, 8)

		for i := 0; i < sizeX; i++ {
			if err := helper.WriteVarint(compressor, int64(d[j][i])); err != nil {
				panic(err.Error())
			}
		}
		compressor.Close()

		encodedImage.Rows[j] = encodedBuffer.Bytes()
	}

	return &encodedImage
}

func (c *LZWCompressor) DecompressLayer(d *data.FileImageLayer) data.Layer {
	sizeY := len(d.Rows)

	image := make([][]float32, sizeY)
	for j := 0; j < sizeY; j++ {
		encodedBuffer := bytes.NewBuffer(d.Rows[j])
		decompressor := lzw.NewReader(encodedBuffer, lzw.MSB, 8)
		decodedSource := bufio.NewReader(decompressor)

		row := []float32{}
		for {
			value, err := binary.ReadVarint(decodedSource)
			if err != nil {
				if err != io.EOF {
					panic(err.Error())
				} else {
					break
				}
			}

			row = append(row, float32(value))
		}
		decompressor.Close()

		image[j] = row
	}

	return image
}
