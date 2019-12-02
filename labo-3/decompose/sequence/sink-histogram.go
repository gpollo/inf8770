package sequence

import (
	"decompose/average"

	"fmt"
	"sync"
)

type HistogramSink struct {
	waitGroup      sync.WaitGroup
	queueHistogram chan *FrameHistogram
	movingAverageA *average.Moving
	movingAverageB *average.Moving
	movingAverage  *average.Moving
}

type averageValue struct {
	value float64
}

func (v *averageValue) Value() float64 {
	return v.value
}

func NewHistogramSink(in chan *FrameHistogram) *HistogramSink {
	return &HistogramSink{
		queueHistogram: in,
		movingAverageA: average.NewMoving(10),
		movingAverageB: average.NewMoving(10),
		movingAverage:  average.NewMoving(10),
	}
}

func (p *HistogramSink) Start() {
	p.waitGroup.Add(1)
	go p.worker()
}

func (p *HistogramSink) Wait() {
	p.waitGroup.Wait()
}

func (p *HistogramSink) worker() {
	defer p.waitGroup.Done()

	var lastFrame *FrameHistogram
	var currentFrame *FrameHistogram
	for {
		nextFrame, ok := <-p.queueHistogram
		if !ok {
			fmt.Println("histogram sink: no more frame...")
			break
		}

		if currentFrame == nil {
			currentFrame = nextFrame
			continue
		}

		if lastFrame == nil {
			lastFrame = currentFrame
			currentFrame = nextFrame
			continue
		}

		a1, a2, a3 := currentFrame.histogram.DifferenceNorm(lastFrame.histogram)
		b1, b2, b3 := currentFrame.histogram.DifferenceNorm(nextFrame.histogram)

		meanA := (a1 + a2 + a3) / 3
		meanB := (b1 + b2 + b3) / 3

		_, filteredA, _ := p.movingAverageA.Add(&averageValue{value: meanA})
		_, filteredB, _ := p.movingAverageB.Add(&averageValue{value: meanB})

		diff := (filteredB - filteredA) / 2
		_, _, _ = p.movingAverage.Add(&averageValue{value: diff})

		lastFrame = currentFrame
		currentFrame = nextFrame
	}
}
