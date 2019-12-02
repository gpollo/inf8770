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

	width := uint(10) / 2
	for {
		edges, ok := <-p.queueEdges
		if !ok {
			fmt.Println("transition pipeline: no more frame...")
			break
		}

		if edges.p < 0.83 {
			p.expected.CheckValidCut(uint(edges.position - 1))
		} else {
			p.expected.CheckInvalidCut(uint(edges.position - 1))

			if 0.88 < edges.p && edges.p < 0.91 {
				p.expected.CheckValidFade(uint(edges.position-1)-width, uint(edges.position-1)+width)
			}
		}
	}

	p.expected.Report()
}
