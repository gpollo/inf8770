package sequence

import (
	"decompose/expected"
	"fmt"
	"sync"
)

type TransitionPipelineSink struct {
	waitGroup  sync.WaitGroup
	queueEdges chan *FrameEdges
	expected   *expected.Expected
}

func NewTransitionPipelineSink(in chan *FrameEdges, e *expected.Expected) *TransitionPipelineSink {
	return &TransitionPipelineSink{
		queueEdges: in,
		expected:   e,
	}
}

func (p *TransitionPipelineSink) Start() {
	p.waitGroup.Add(1)
	go p.worker()
}

func (p *TransitionPipelineSink) Wait() {
	p.waitGroup.Wait()
}

func (p *TransitionPipelineSink) worker() {
	defer p.waitGroup.Done()

	for {
		edges, ok := <-p.queueEdges
		if !ok {
			fmt.Println("transition pipeline: no more frame...")
			break
		}

		fmt.Println("DATA", edges.position-1, edges.pOut, edges.pIn, edges.p)

		if edges.p < 0.88 {
			//p.expected.CheckValidCut(uint(edges.position - 1))
			//fmt.Println(edges.position-1, edges.pOut, edges.pIn, edges.p)
		} else {
			//p.expected.CheckInvalidCut(uint(edges.position - 1))
		}
	}
}
