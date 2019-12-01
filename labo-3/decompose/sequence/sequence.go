package sequence

import (
	"decompose/edges"
	"decompose/expected"
	"decompose/helper"
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

func (s *Sequence) Run(e *expected.Expected, save bool) {
	bufferSize := 500
	queueFilenames := make(chan *FrameFilename, bufferSize)
	queueImages := make(chan *FrameImage, bufferSize)
	queueUnorderedSobel := make(chan *FrameSobel, bufferSize)
	queueOrderedSobel := make(chan *FrameSobel, bufferSize)
	queueSobelPair := make(chan *FrameSobelPair, bufferSize)
	queueUnorderedEdges := make(chan *FrameEdges, bufferSize)
	queueOrderedEdges := make(chan *FrameEdges, bufferSize)
	filter := sobel.With33Kernel()
	read := NewReadFilePipeline(queueFilenames, queueImages, 3)
	sobel := NewSobelPipeline(queueImages, queueUnorderedSobel, 20, filter)
	sobelOrderer := NewSobelOrdererPipeline(queueUnorderedSobel, queueOrderedSobel)
	sobelPair := NewSobelPairPipeline(queueOrderedSobel, queueSobelPair)
	edges := NewEdgesPipeline(queueSobelPair, queueUnorderedEdges, 10)
	edgesOrderer := NewEdgesOrdererPipeline(queueUnorderedEdges, queueOrderedEdges)
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

	position := 0
	for {
		position += 1
		filename := s.inputDirectory + "/" + fmt.Sprintf(s.inputExpr, position)

		_, err := os.Stat(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			break
		}

		queueFilenames <- &FrameFilename{
			position: position,
			filename: filename,
		}
	}
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

func (s *Sequence) saveTempImage(position int, i image.Image) error {
	filename := s.tempDirectory + fmt.Sprintf(s.tempExpr, position)
	if err := helper.SaveImage(i, filename); err != nil {
		return err
	}

	return nil
}
