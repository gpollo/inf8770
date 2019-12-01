package sequence

import (
	"fmt"
	"math"
	"sync"
)

type EdgesPipeline struct {
	workerCount    int
	waitGroup      sync.WaitGroup
	queueSobelPair chan *FrameSobelPair
	queueEdges     chan *FrameEdges
}

func NewEdgesPipeline(in chan *FrameSobelPair, out chan *FrameEdges, threads uint) *EdgesPipeline {
	return &EdgesPipeline{
		workerCount:    int(threads),
		queueSobelPair: in,
		queueEdges:     out,
	}
}

func (p *EdgesPipeline) Start() {
	p.waitGroup.Add(p.workerCount)
	for i := 0; i < p.workerCount; i++ {
		go p.worker()
	}
}

func (p *EdgesPipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queueEdges)
}

func (p *EdgesPipeline) worker() {
	defer p.waitGroup.Done()

	for {
		pair, ok := <-p.queueSobelPair
		if !ok {
			fmt.Println("edges pipeline: no more sobel frame...")
			break
		}

		entering := pair.edgesExpanded.Contains(pair.nextFrame.edgesInitial)
		exiting := pair.nextFrame.edgesExpanded.Contains(pair.edgesInitial)

		pIn := 1.0 - float64(entering.Count())/float64(pair.nextFrame.edgesInitial.Count())
		pOut := 1.0 - float64(exiting.Count())/float64(pair.edgesInitial.Count())
		pMax := math.Min(pIn, pOut)

		p.queueEdges <- &FrameEdges{
			FrameSobelPair: *pair,
			edgesEntering:  entering,
			edgesExiting:   exiting,
			pOut:           pOut,
			pIn:            pIn,
			p:              pMax,
		}
	}
}
