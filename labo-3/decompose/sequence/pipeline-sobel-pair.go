package sequence

import (
	"fmt"
	"sync"
)

type SobelPairPipeline struct {
	waitGroup  sync.WaitGroup
	queueSobel chan *FrameSobel
	queuePair  chan *FrameSobelPair
}

func NewSobelPairPipeline(in chan *FrameSobel, out chan *FrameSobelPair) *SobelPairPipeline {
	return &SobelPairPipeline{
		queueSobel: in,
		queuePair:  out,
	}
}

func (p *SobelPairPipeline) Start() {
	p.waitGroup.Add(1)
	go p.worker()
}

func (p *SobelPairPipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queuePair)
}

func (p *SobelPairPipeline) worker() {
	defer p.waitGroup.Done()

	var current *FrameSobel = nil
	for {
		next, ok := <-p.queueSobel
		if !ok {
			fmt.Println("sobel pair pipeline: no more frame...")
			break
		}

		if current == nil {
			current = next
			continue
		}

		p.queuePair <- &FrameSobelPair{
			FrameSobel: *current,
			nextFrame:  next,
		}

		current = next
	}
}
