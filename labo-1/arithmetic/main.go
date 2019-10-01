package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"time"

	"github.com/akamensky/argparse"
)

const SYMBOL_COUNT = 256

var verbose *bool
var parallel *bool
var workers *int

func main() {
	parser := argparse.NewParser("encode", "STDIN arithmetic encoder/decoder")
	doDecode := parser.Flag("d", "decode", &argparse.Options{
		Required: false,
		Help:     "Run the decoder on the input data.",
	})
	doEncode := parser.Flag("e", "encode", &argparse.Options{
		Required: false,
		Help:     "Run the encoder on the input data. [default]",
	})
	doBenchmark := parser.Flag("b", "benchmark", &argparse.Options{
		Required: false,
		Help:     "Print execution time (ms) and don't output data",
	})

	parallel = parser.Flag("p", "parallel", &argparse.Options{
		Required: false,
		Help:     "Run interval resizing using multiple threads",
	})
	workers = parser.Int("w", "workers", &argparse.Options{
		Required: false,
		Help:     "Number of workers when `--parallel` arguments is used",
		Default:  4,
	})
	verbose = parser.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Print debug informations while encoding.",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if !*doEncode && !*doDecode {
		*doEncode = true
	}

	if *doEncode && *doDecode {
		fmt.Print("Cannot have both `--encode` and `--decode` at the same time")
		os.Exit(1)
	}

	var stdin *bufio.Reader
	var stdout *bufio.Writer

	stdin = bufio.NewReader(os.Stdin)
	if *doBenchmark {
		stdout = bufio.NewWriter(ioutil.Discard)
	} else {
		stdout = bufio.NewWriter(os.Stdout)
	}

	defer stdout.Flush()

	startTimestamp := time.Now().UnixNano()
	if *doEncode {
		if err := Encode(stdin, stdout); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	} else {
		if err := Decode(stdin, stdout); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
	stopTimestamp := time.Now().UnixNano()

	if *doBenchmark {
		elapsed := stopTimestamp - startTimestamp
		fmt.Printf("%d us\n", elapsed/1000000)
	}
}
