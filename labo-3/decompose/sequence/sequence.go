package sequence

import (
	"decompose/edges"
	"decompose/helper"
	"decompose/histogram"
	"decompose/layer"
	"decompose/sobel"
	"fmt"
	"image"
	"os"
)

type Sequence struct {
	frameCount    uint
	histogramBins uint
	histogramLast *histogram.Histogram
}

func FromDirectory(directory, expr string) (*Sequence, error) {
	s := Sequence{
		frameCount:    0,
		histogramBins: 100,
		histogramLast: nil,
	}

	for frame := 1; true; frame++ {
		filename := directory + "/" + fmt.Sprintf(expr, frame)
		_, err := os.Stat(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			break
		}

		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		image, _, err := image.Decode(file)
		if err != nil {
			return nil, err
		}

		filter := sobel.With55Kernel()

		layerBW := layer.FromImageBW(image)
		layerSobel := filter.Apply(layerBW)
		layerEdges := edges.FromLayer(layerSobel, 35000)
		layerExpanded := layerEdges.ExpandRadius(5)
		filename = "tmp/" + fmt.Sprintf(expr, frame)
		if err = helper.SaveImage(layerExpanded, filename); err != nil {
			return nil, err
		}

		s.AddFrame(image)
		file.Close()
	}

	return &s, nil
}

func (s *Sequence) AddFrame(i image.Image) {
	s.frameCount++
	if s.histogramLast == nil {
		s.histogramLast = histogram.FromImage(i, s.histogramBins)
		return
	}

	histogramThis := histogram.FromImage(i, s.histogramBins)
	fmt.Println(s.frameCount)
	fmt.Println(histogramThis.DifferenceNorm(s.histogramLast))
	s.histogramLast = histogramThis
}
