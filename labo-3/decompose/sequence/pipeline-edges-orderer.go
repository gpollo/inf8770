package sequence

import (
	"fmt"
	"sync"
)

type EdgesOrdererPipeline struct {
	waitGroup      sync.WaitGroup
	queueUnordered chan *FrameEdges
	queueOrdered   chan *FrameEdges
	received       map[int]*FrameEdges
	position       int
}

func NewEdgesOrdererPipeline(in, out chan *FrameEdges) *EdgesOrdererPipeline {
	return &EdgesOrdererPipeline{
		queueUnordered: in,
		queueOrdered:   out,
		received:       make(map[int]*FrameEdges),
		position:       1,
	}
}

func (p *EdgesOrdererPipeline) Start() {
	p.waitGroup.Add(1)
	go p.worker()
}

func (p *EdgesOrdererPipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queueOrdered)
}

func (p *EdgesOrdererPipeline) worker() {
	defer p.waitGroup.Done()

	for {
		unordered, ok := <-p.queueUnordered
		if !ok {
			fmt.Println("edges orderer pipeline: no more frame...")
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
			p.position += 1
		}
	}
}
