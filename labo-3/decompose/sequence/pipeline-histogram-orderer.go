package sequence

import (
	"fmt"
	"sync"
)

type HistogramOrdererPipeline struct {
	waitGroup      sync.WaitGroup
	queueUnordered chan *FrameHistogram
	queueOrdered   chan *FrameHistogram
	received       map[int]*FrameHistogram
	position       int
	skip           int
}

func NewHistogramOrdererPipeline(in, out chan *FrameHistogram, skip int) *HistogramOrdererPipeline {
	return &HistogramOrdererPipeline{
		queueUnordered: in,
		queueOrdered:   out,
		received:       make(map[int]*FrameHistogram),
		position:       1,
		skip:           skip,
	}
}

func (p *HistogramOrdererPipeline) Start() {
	p.waitGroup.Add(1)
	go p.worker()
}

func (p *HistogramOrdererPipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queueOrdered)
}

func (p *HistogramOrdererPipeline) worker() {
	defer p.waitGroup.Done()

	for {
		unordered, ok := <-p.queueUnordered
		if !ok {
			fmt.Println("histogram orderer pipeline: no more frame...")
			break
		}

		p.received[unordered.position] = unordered

		for {
			ordered, ok := p.received[p.position]
			if !ok {
				break
			}
			p.queueOrdered <- ordered

			delete(p.received, p.position)
			p.position += p.skip
		}
	}
}
