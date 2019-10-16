package wavelet

import (
	"jpeg2000/data"
	"jpeg2000/helper"

	proto "github.com/golang/protobuf/proto"
	"os"
	"path"
)

type PyWavelet struct {
	mode string
}

func getPythonScript(script string) string {
	if directory, ok := os.LookupEnv("JPEG2000_PYTHON"); ok {
		return path.Join(directory, script)
	} else {
		return path.Join("../python/", script)
	}
}

func (w *PyWavelet) WaveletTransform(l data.Layer) data.Layer {
	image := l.ToProtobuf()
	dwt := data.PythonDWT{Mode: w.mode, Data: &image}
	pdata, err := proto.Marshal(&dwt)
	if err != nil {
		panic(err.Error())
	}

	args := []string{"python3", getPythonScript("dwt.py")}
	result, err := helper.CallProcess(args, pdata)
	if err != nil {
		panic(err.Error())
	}

	wavelet := data.ImageData{}
	err = proto.Unmarshal(result, &wavelet)
	if err != nil {
		panic(err.Error())
	}

	decoded := data.Layer{}
	decoded.FromProtobuf(wavelet)
	decoded.Times(0.5)

	return decoded
}

func (w *PyWavelet) WaveletInverse(l data.Layer) data.Layer {
	image := l.ToProtobuf()
	dwt := data.PythonDWT{Mode: w.mode, Data: &image}
	pdata, err := proto.Marshal(&dwt)
	if err != nil {
		panic(err.Error())
	}

	args := []string{"python3", getPythonScript("idwt.py")}
	result, err := helper.CallProcess(args, pdata)
	if err != nil {
		panic(err.Error())
	}

	wavelet := data.ImageData{}
	err = proto.Unmarshal(result, &wavelet)
	if err != nil {
		panic(err.Error())
	}

	decoded := data.Layer{}
	decoded.FromProtobuf(wavelet)
	decoded.Times(2.0)

	return decoded
}
