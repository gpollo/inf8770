package main

import (
	"image"
	"jpeg2000/compressor"
	"jpeg2000/data"
	"jpeg2000/quantifier"
	"jpeg2000/subsampler"
	"jpeg2000/wavelet"

	proto "github.com/golang/protobuf/proto"
)

type Pipeline struct {
	conversion bool
	subsampler subsampler.Subsampler
	wavelet    wavelet.Wavelet
	quantifier quantifier.Quantifier
	compressor compressor.Compressor
}

func (p *Pipeline) GetProtobufHeader(w, h uint) *data.FileImageHeader {
	return &data.FileImageHeader{
		Width:       uint32(w),
		Height:      uint32(h),
		Conversion:  p.conversion,
		Subsampling: p.subsampler.ToProtobuf(),
		Wavelet:     p.wavelet.ToProtobuf(),
		Quantifier:  p.quantifier.ToProtobuf(),
	}
}

func (p *Pipeline) EncodeImage(input image.Image) ([]byte, error) {
	var err error

	r, g, b, w, h := data.GetLayers(input)
	rgb := data.Image{r, g, b}

	yuv := data.Image{}
	if p.conversion {
		yuv, err = rgb.RGBToYUV()
		if err != nil {
			return []byte{}, err
		}
	} else {
		yuv = rgb
	}
	ys, us, vs := yuv[0], yuv[1], yuv[2]

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

	p.conversion = d.Conversion

	p.subsampler, err = subsampler.FromProtobuf(d.Subsampling)
	if err != nil {
		return err
	}

	p.wavelet, err = wavelet.FromProtobuf(d.Wavelet)
	if err != nil {
		return err
	}

	p.quantifier, err = quantifier.FromProtobuf(d.Quantifier)
	if err != nil {
		return err
	}

	p.compressor = &compressor.LZWCompressor{}

	return nil
}

func (p *Pipeline) DecodeImage(input []byte) (image.Image, error) {
	var err error

	encoded := &data.FileImage{}
	if err = proto.Unmarshal(input, encoded); err != nil {
		return nil, err
	}

	if err = p.SetupFromProtobufHeader(encoded.Header); err != nil {
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
	yuv := data.Image{ys, us, vs}

	rgb := data.Image{}
	if p.conversion {
		rgb, err = yuv.YUVToRGB()
		if err != nil {
			return nil, err
		}
	} else {
		rgb = yuv
	}

	return &rgb, nil
}
