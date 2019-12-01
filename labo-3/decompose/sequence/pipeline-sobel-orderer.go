package sequence

import (
	"fmt"
	"sync"
)

type SobelOrdererPipeline struct {
	waitGroup      sync.WaitGroup
	queueUnordered chan *FrameSobel
	queueOrdered   chan *FrameSobel
	received       map[int]*FrameSobel
	position       int
}

func NewSobelOrdererPipeline(in, out chan *FrameSobel) *SobelOrdererPipeline {
	return &SobelOrdererPipeline{
		queueUnordered: in,
		queueOrdered:   out,
		received:       make(map[int]*FrameSobel),
		position:       1,
	}
}

func (p *SobelOrdererPipeline) Start() {
	p.waitGroup.Add(1)
	go p.worker()
}

func (p *SobelOrdererPipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queueOrdered)
}

func (p *SobelOrdererPipeline) worker() {
	defer p.waitGroup.Done()

	for {
		unordered, ok := <-p.queueUnordered
		if !ok {
			fmt.Println("sobel orderer pipeline: no more frame...")
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
