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
	r, g, b, w, h := data.GetLayers(input)
	ys, us, vs := data.RGBToYUV(r, g, b)
	y, u, v := p.subsampler.Subsample(ys, us, vs)

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
		Data:   &data.FileImageData{Y: yc, U: uc, V: vc},
	}

	pdata, err := proto.Marshal(&encoded)
	if err != nil {
		return nil, err
	}

	return pdata, nil
}

func (p *Pipeline) SetupFromProtobufHeader(d *data.FileImageHeader) error {
	var err error

	p.subsampler, err = SubsamplerFromProtobuf(d.Subsampling)
	if err != nil {
		return err
	}

	p.wavelet, err = WaveletFromProtobuf(d.Wavelet)
	if err != nil {
		return err
	}

	p.quantifier, err = QuantifierFromProtobuf(d.Quantifier)
	if err != nil {
		return err
	}

	p.compressor = &LZWCompressor{}

	return nil
}

func (p *Pipeline) DecodeImage(input []byte) (image.Image, error) {
	encoded := &data.FileImage{}
	if err := proto.Unmarshal(input, encoded); err != nil {
		return nil, err
	}

	if err := p.SetupFromProtobufHeader(encoded.Header); err != nil {
		return nil, err
	}

	yc := encoded.Data.Y
	uc := encoded.Data.U
	vc := encoded.Data.V

	yq := p.compressor.DecompressLayer(yc)
	uq := p.compressor.DecompressLayer(uc)
	vq := p.compressor.DecompressLayer(vc)

	yw := p.quantifier.QuantifierInverse(yq)
	uw := p.quantifier.QuantifierInverse(uq)
	vw := p.quantifier.QuantifierInverse(vq)

	y := p.wavelet.WaveletInverse(yw)
	u := p.wavelet.WaveletInverse(uw)
	v := p.wavelet.WaveletInverse(vw)

	ys, us, vs := p.subsampler.Supersample(y, u, v)
	r, g, b := data.YUVToRGB(ys, us, vs)

	return &Image{r, g, b}, nil
}
