package main

import (
	"image"
	"jpeg2000/data"

	proto "github.com/golang/protobuf/proto"
)

type Pipeline struct {
	subsampler Subsampler
	wavelet    Wavelet
	quantifier Quantifier
	compressor Compressor
}

func (p *Pipeline) GetProtobufHeader(w, h uint) *data.FileImageHeader {
	return &data.FileImageHeader{
		Width:       uint32(w),
		Height:      uint32(h),
		Subsampling: p.subsampler.ToProtobuf(),
		Wavelet:     p.wavelet.ToProtobuf(),
		Quantifier:  p.quantifier.ToProtobuf(),
	}
}

func (p *Pipeline) EncodeImage(input image.Image) ([]byte, error) {
	r, g, b, w, h := GetImageData(input)
	y, u, v := p.subsampler.Subsample(r, g, b)

	yw := p.wavelet.WaveletTransform(y)
	uw := p.wavelet.WaveletTransform(u)
	vw := p.wavelet.WaveletTransform(v)

	yq := p.quantifier.QuantifierTransform(yw)
	uq := p.quantifier.QuantifierTransform(uw)
	vq := p.quantifier.QuantifierTransform(vw)

	yc := p.compressor.CompressLayer(yq)
	uc := p.compressor.CompressLayer(uq)
	vc := p.compressor.CompressLayer(vq)

	encoded := data.FileImage{
		Header: p.GetProtobufHeader(w, h),
		Data:   &data.FileImageData{Y: &yc, U: &uc, V: &vc},
	}

	pdata, err := proto.Marshal(&encoded)
	if err != nil {
		return nil, err
	}

	return pdata, nil
}

func DecodeImage(input []byte) (image.Image, error) {
	return nil, nil
}
