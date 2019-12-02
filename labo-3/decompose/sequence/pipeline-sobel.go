package sequence

import (
	"decompose/edges"
	"decompose/layer"
	"decompose/sobel"
	"fmt"
	"sync"
)

type SobelPipeline struct {
	workerCount int
	waitGroup   sync.WaitGroup
	queueImages chan *FrameImage
	queueSobel  chan *FrameSobel
	filter      *sobel.Sobel
}

func NewSobelPipeline(in chan *FrameImage, out chan *FrameSobel, threads uint, filter *sobel.Sobel) *SobelPipeline {
	return &SobelPipeline{
		workerCount: int(threads),
		queueImages: in,
		queueSobel:  out,
		filter:      filter,
	}
}

func (p *SobelPipeline) Start() {
	p.waitGroup.Add(p.workerCount)
	for i := 0; i < p.workerCount; i++ {
		go p.worker()
	}
}

func (p *SobelPipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queueSobel)
}

func (p *SobelPipeline) worker() {
	defer p.waitGroup.Done()

	for {
		image, ok := <-p.queueImages
		if !ok {
			fmt.Println("sobel pipeline: no more image...")
			break
		}

		layerBW := layer.FromImageBW(image.image)
		layerSobel := p.filter.Apply(layerBW)
		layerEdges := edges.FromLayer(layerSobel, 40000)
		layerExpanded := layerEdges.ExpandRadius(7)

		p.queueSobel <- &FrameSobel{
			FrameImage:    *image,
			layerBW:       layerBW,
			layerSobel:    layerSobel,
			edgesInitial:  layerEdges,
			edgesExpanded: layerExpanded,
		}
	}
}
