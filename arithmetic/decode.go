package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"time"
)

type Decoder struct {
	total     int
	symbols   [SYMBOL_COUNT]Symbol
	intervals Intervals
	float     big.Rat
}

func (e *Decoder) getSymbols(input *bufio.Reader) error {
	for i := 0; i < SYMBOL_COUNT; i++ {
		e.symbols[i].number = 0
		e.symbols[i].count = 0
		e.symbols[i].frequency.SetInt64(0)
	}

	data, err := input.Peek(5)
	if err != nil {
		return err
	}

	e.total = 0
	e.total += (int(data[0]) << 0)
	e.total += (int(data[1]) << 8)
	e.total += (int(data[2]) << 16)
	e.total += (int(data[3]) << 24)

	size := int(data[4])

	_, err = input.Discard(5)
	if err != nil {
		return err
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "header: total=%d\n", e.total)
		fmt.Fprintf(os.Stderr, "header: symbol count=%d\n", size)
	}

	for {
		data, err = input.Peek(5)
		if err != nil {
			return err
		}

		number := byte(data[0])

		count := 0
		count += (int(data[1]) << 0)
		count += (int(data[2]) << 8)
		count += (int(data[3]) << 16)
		count += (int(data[4]) << 24)

		_, err = input.Discard(5)
		if err != nil {
			return err
		}

		if *debug {
			fmt.Fprintf(os.Stderr, "header: symbol=%d (%c), count=%d\n",
				number, number, count,
			)
		}

		e.symbols[number].number = number
		e.symbols[number].count = count
		e.symbols[number].computeFrequency(e.total)

		size--
		if size == 0 {
			break
		}
	}

	return nil
}

func (e *Decoder) GetFloat(input *bufio.Reader) error {
	base := big.NewInt(2)
	exponent := big.NewInt(0)
	increment := big.NewInt(1)
	power := big.NewInt(0)
	frac := big.NewRat(0, 1)

	e.float.SetInt64(0)

	data := make([]byte, 1024)
	for {
		n, err := input.Read(data)
		if n == 0 {
			break
		}

		if err != nil {
			return err
		}

		for i := 0; i < n; i++ {
			for j := 0; j < 8; j++ {
				exponent.Add(exponent, increment)

				if data[i]&(1<<j) != 0 {
					power.Exp(base, exponent, nil)
					frac.SetFrac(increment, power)
					e.float.Add(&e.float, frac)
				}
			}
		}
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "Initial float: %s\n", e.float.FloatString(50))
	}

	return nil
}

func (e *Decoder) getIntervals() {
	sorted := Frequencies{}
	for i := 0; i < SYMBOL_COUNT; i++ {
		if e.symbols[i].count == 0 {
			continue
		}

		sorted = append(sorted, e.symbols[i])
	}
	sort.Sort(sorted)

	for i := 0; i < len(sorted); i++ {
		e.intervals.Append(sorted[i].number, &sorted[i].frequency)
	}

	if *debug {
		e.intervals.Print()
	}
}

func (e *Decoder) Decode(output *bufio.Writer) error {
	for i := 0; i < e.total; i++ {
		var tsBegin, tsEnd int64
		if *timed {
			tsBegin = time.Now().UnixNano()
		}

		interval, err := e.intervals.FindByRat(&e.float)
		if err != nil {
			return err
		}

		e.intervals.Resize(interval)

		if *timed {
			tsEnd = time.Now().UnixNano()
			fmt.Fprintf(os.Stderr, "%d,%d\n", i, tsEnd-tsBegin)
		}

		output.WriteByte(interval.number)
	}

	return nil
}

func decodeData(in *bufio.Reader, out *bufio.Writer) error {
	decoder := Decoder{}

	if err := decoder.getSymbols(in); err != nil {
		return err
	}

	if err := decoder.GetFloat(in); err != nil {
		return err
	}

	decoder.getIntervals()

	if err := decoder.Decode(out); err != nil {
		return err
	}

	return nil
}
