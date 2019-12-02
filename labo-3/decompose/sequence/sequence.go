package sequence

import (
	"decompose/edges"
	"decompose/expected"
	"decompose/histogram"
	"decompose/layer"
	"decompose/sobel"
	"fmt"
	"image"
	"os"
)

type FrameFilename struct {
	position int
	filename string
}

type FrameImage struct {
	FrameFilename
	image image.Image
}

type FrameSobel struct {
	FrameImage
	layerBW       layer.Layer
	layerSobel    layer.Layer
	edgesInitial  edges.Edges
	edgesExpanded edges.Edges
}

type FrameSobelPair struct {
	FrameSobel
	nextFrame *FrameSobel
}

type FrameEdges struct {
	FrameSobelPair
	edgesEntering edges.Edges
	edgesExiting  edges.Edges
	pOut          float64
	pIn           float64
	p             float64
}

type FrameHistogram struct {
	FrameImage
	histogram *histogram.Histogram
}

type FrameDistance struct {
	FrameHistogram
	distance float64
}

type Sequence struct {
	WorkerCount    int
	inputDirectory string
	inputExpr      string
	tempDirectory  string
	tempExpr       string
}

func FromDirectory(directory, expr string) *Sequence {
	return &Sequence{
		WorkerCount:    200,
		inputDirectory: directory,
		inputExpr:      expr,
		tempDirectory:  "tmp/",
		tempExpr:       expr,
	}
}

func (s *Sequence) RunSobel(e *expected.Expected, save bool, skip int) {
	bufferSize := 500
	queueFilenames := make(chan *FrameFilename, bufferSize)
	queueImages := make(chan *FrameImage, bufferSize)
	queueUnorderedSobel := make(chan *FrameSobel, bufferSize)
	queueOrderedSobel := make(chan *FrameSobel, bufferSize)
	queueSobelPair := make(chan *FrameSobelPair, bufferSize)
	queueUnorderedEdges := make(chan *FrameEdges, bufferSize)
	queueOrderedEdges := make(chan *FrameEdges, bufferSize)
	filter := sobel.With55Kernel()
	read := NewReadFilePipeline(queueFilenames, queueImages, 3)
	sobel := NewSobelPipeline(queueImages, queueUnorderedSobel, 20, filter)
	sobelOrderer := NewSobelOrdererPipeline(queueUnorderedSobel, queueOrderedSobel, skip)
	sobelPair := NewSobelPairPipeline(queueOrderedSobel, queueSobelPair)
	edges := NewEdgesPipeline(queueSobelPair, queueUnorderedEdges, 10)
	edgesOrderer := NewEdgesOrdererPipeline(queueUnorderedEdges, queueOrderedEdges, skip)
	sinkSave := NewEdgesSavePipelineSink(queueOrderedEdges, 6, s.tempDirectory+s.tempExpr)
	sinkFind := NewTransitionPipelineSink(queueOrderedEdges, e)

	read.Start()
	sobel.Start()
	sobelOrderer.Start()
	sobelPair.Start()
	edges.Start()
	edgesOrderer.Start()
	if save {
		sinkSave.Start()
	} else {
		sinkFind.Start()
	}

	s.queueFilenames(queueFilenames, skip)
	close(queueFilenames)

	read.Wait()
	sobel.Wait()
	sobelOrderer.Wait()
	sobelPair.Wait()
	edges.Wait()
	edgesOrderer.Wait()
	if save {
		sinkSave.Wait()
	} else {
		sinkFind.Wait()
	}
}

func (s *Sequence) RunHistogram(e *expected.Expected, bins uint, skip int) {
	bufferSize := 500
	queueFilenames := make(chan *FrameFilename, bufferSize)
	queueImages := make(chan *FrameImage, bufferSize)
	queueUnorderedHistogram := make(chan *FrameHistogram, bufferSize)
	queueOrderedHistogram := make(chan *FrameHistogram, bufferSize)

	read := NewReadFilePipeline(queueFilenames, queueImages, 5)
	histogram := NewHistogramPipeline(queueImages, queueUnorderedHistogram, 20, bins)
	histogramOrderer := NewHistogramOrdererPipeline(queueUnorderedHistogram, queueOrderedHistogram, skip)
	histogramSink := NewHistogramSink(queueOrderedHistogram)

	read.Start()
	histogram.Start()
	histogramOrderer.Start()
	histogramSink.Start()

	s.queueFilenames(queueFilenames, skip)
	close(queueFilenames)

	read.Wait()
	histogram.Wait()
	histogramOrderer.Wait()
	histogramSink.Wait()
}

func (s *Sequence) queueFilenames(queue chan *FrameFilename, skip int) {
	position := 1
	for {
		filename := s.inputDirectory + "/" + fmt.Sprintf(s.inputExpr, position)

		_, err := os.Stat(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			break
		}

		queue <- &FrameFilename{
			position: position,
			filename: filename,
		}

		position += skip
	}
}
