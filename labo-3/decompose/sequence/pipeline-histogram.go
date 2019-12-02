package sequence

import (
	"decompose/histogram"
	"fmt"
	"sync"
)

type HistogramPipeline struct {
	workerCount    int
	waitGroup      sync.WaitGroup
	queueImages    chan *FrameImage
	queueHistogram chan *FrameHistogram
	bins           uint
}

func NewHistogramPipeline(in chan *FrameImage, out chan *FrameHistogram, threads uint, bins uint) *HistogramPipeline {
	return &HistogramPipeline{
		workerCount:    int(threads),
		queueImages:    in,
		queueHistogram: out,
		bins:           bins,
	}
}

func (p *HistogramPipeline) Start() {
	p.waitGroup.Add(p.workerCount)
	for i := 0; i < p.workerCount; i++ {
		go p.worker()
	}
}

func (p *HistogramPipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queueHistogram)
}

func (p *HistogramPipeline) worker() {
	defer p.waitGroup.Done()

	for {
		image, ok := <-p.queueImages
		if !ok {
			fmt.Println("histogram pipeline: no more image...")
			break
		}

		p.queueHistogram <- &FrameHistogram{
			FrameImage: *image,
			histogram:  histogram.FromImage(image.image, p.bins),
		}
	}
}
