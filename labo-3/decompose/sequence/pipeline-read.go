package sequence

import (
	"fmt"
	"image"
	"os"
	"sync"
)

type ReadFilePipeline struct {
	workerCount    int
	waitGroup      sync.WaitGroup
	queueFilenames chan *FrameFilename
	queueImages    chan *FrameImage
}

func NewReadFilePipeline(in chan *FrameFilename, out chan *FrameImage, threads uint) *ReadFilePipeline {
	return &ReadFilePipeline{
		workerCount:    int(threads),
		queueFilenames: in,
		queueImages:    out,
	}
}

func (p *ReadFilePipeline) Start() {
	p.waitGroup.Add(p.workerCount)
	for i := 0; i < p.workerCount; i++ {
		go p.worker()
	}
}

func (p *ReadFilePipeline) Wait() {
	p.waitGroup.Wait()
	close(p.queueImages)
}

func (p *ReadFilePipeline) worker() {
	defer p.waitGroup.Done()

	for {
		filename, ok := <-p.queueFilenames
		if !ok {
			fmt.Println("read pipeline: no more filename...")
			break
		}

		file, err := os.Open(filename.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read pipeline: %s\n", err.Error())
			break
		}

		image, _, err := image.Decode(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read pipeline: %s\n", err.Error())
			break
		}
		file.Close()

		p.queueImages <- &FrameImage{
			FrameFilename: *filename,
			image:         image,
		}
	}
}
