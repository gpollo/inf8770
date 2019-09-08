package main

import "fmt"
import "bufio"
import "math/big"
import "os"
import "sort"

const SYMBOL_COUNT = 256

var symbols [SYMBOL_COUNT]Symbol
var intervals Intervals
var total uint = 0
var data []byte

func countSymbols() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		b, err := stdin.ReadByte()
		if err != nil {
			break
		}

		symbols[b].number = b
		symbols[b].count++
		total++

		data = append(data, b)
	}
}

func computeFrequencies() {
	for i := 0; i < SYMBOL_COUNT; i++ {
		symbols[i].computeFrequency(total)
	}
}

func buildIntervals() {
	sorted := Frequencies{}
	for i := 0; i < SYMBOL_COUNT; i++ {
		if symbols[i].count == 0 {
			continue
		}

		sorted = append(sorted, symbols[i])
	}
	sort.Sort(sorted)

	current := big.NewFloat(0)
	for i := 0; i < len(sorted); i++ {
		symbol := sorted[i]
		interval := Interval{}
		interval.number = symbol.number
		interval.start.Copy(current)
		current.Add(current, &symbol.frequency)
		interval.end.Copy(current)
		intervals = append(intervals, interval)
	}
}

func encodeSymbols() error {
	for i := 0; i < len(data); i++ {
		symbol := data[i]

		interval, err := intervals.find(symbol)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("------------------------------")
		printIntervals()
		fmt.Fprintf(os.Stderr, "interval ayant %d: %s to %s\n",
			interval.number,
			interval.start.Text('f', 30),
			interval.end.Text('f', 30),
		)

		intervals.resize(interval)
	}

	return nil
}

func printSymbols() {
	for i := 0; i < SYMBOL_COUNT; i++ {
		if symbols[i].count == 0 {
			continue
		}

		fmt.Fprintf(os.Stderr, "%02x (%c): %d %s\n",
			i,
			symbols[i].number,
			symbols[i].count,
			symbols[i].frequency.Text('f', 20),
		)
	}
}

func printIntervals() {
	for i := 0; i < len(intervals); i++ {
		fmt.Fprintf(os.Stderr, "%02x (%c): %s to %s\n",
			intervals[i].number,
			intervals[i].number,
			intervals[i].start.Text('f', 50),
			intervals[i].end.Text('f', 50),
		)
	}
}

func main() {
	countSymbols()
	computeFrequencies()
	buildIntervals()
	encodeSymbols()
}
