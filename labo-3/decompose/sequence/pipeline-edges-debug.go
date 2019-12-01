package sequence

import (
	"decompose/edges"
	"decompose/helper"
	"fmt"
	"os"
	"sync"
)

type EdgesSavePipelineSink struct {
	workerCount int
	waitGroup   sync.WaitGroup
	queueEdges  chan *FrameEdges
	saveExpr    string
}

func NewEdgesSavePipelineSink(in chan *FrameEdges, threads uint, expr string) *EdgesSavePipelineSink {
	return &EdgesSavePipelineSink{
		workerCount: int(threads),
		queueEdges:  in,
		saveExpr:    expr,
	}
}

func (p *EdgesSavePipelineSink) Start() {
	p.waitGroup.Add(p.workerCount)
	for i := 0; i < p.workerCount; i++ {
		go p.worker()
	}
}

func (p *EdgesSavePipelineSink) Wait() {
	p.waitGroup.Wait()
}

func (p *EdgesSavePipelineSink) worker() {
	defer p.waitGroup.Done()

	for {
		current, ok := <-p.queueEdges
		if !ok {
			fmt.Println("edges save pipeline: no more frame...")
			break
		}

		merged := edges.From6Edges(
			current.edgesInitial,
			current.nextFrame.edgesInitial,
			current.edgesExpanded,
			current.nextFrame.edgesExpanded,
			current.edgesEntering,
			current.edgesExiting,
		)

		filename := fmt.Sprintf(p.saveExpr, current.position)
		if err := helper.SaveImage(merged, filename); err != nil {
			fmt.Fprintf(os.Stderr, "edges save pipeline: %s\n", err.Error())
			continue
		}
		fmt.Printf("edges save pipeline: frame #%d\n", current.position)
	}
}
