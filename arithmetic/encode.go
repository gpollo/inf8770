package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"time"
)

type Encoder struct {
	data      []byte
	symbols   [SYMBOL_COUNT]Symbol
	intervals Intervals
}

func (e *Encoder) getSymbols(input *bufio.Reader, output *bufio.Writer) error {
	for {
		b, err := input.ReadByte()
		if err != nil {
			break
		}

		e.data = append(e.data, b)
	}

	for i := 0; i < len(e.symbols); i++ {
		e.symbols[i].number = byte(i)
		e.symbols[i].count = 0
		e.symbols[i].frequency.SetInt64(0)
	}

	for i := 0; i < len(e.data); i++ {
		e.symbols[e.data[i]].count++
	}

	for i := 0; i < len(e.symbols); i++ {
		e.symbols[i].computeFrequency(len(e.data))
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "header: total=%d\n", len(e.data))
	}

	_, err := output.Write([]byte{
		byte(0xFF & (len(e.data) >> 0)),
		byte(0xFF & (len(e.data) >> 8)),
		byte(0xFF & (len(e.data) >> 16)),
		byte(0xFF & (len(e.data) >> 24)),
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *Encoder) getIntervals(output *bufio.Writer) error {
	sorted := Frequencies{}
	for i := 0; i < len(e.symbols); i++ {
		if e.symbols[i].count == 0 {
			continue
		}

		sorted = append(sorted, e.symbols[i])
	}
	sort.Sort(sorted)

	if *debug {
		fmt.Fprintf(os.Stderr, "header: symbol count=%d\n", len(sorted))
	}

	_, err := output.Write([]byte{
		byte(len(sorted)),
	})
	if err != nil {
		return err
	}

	for i := 0; i < len(sorted); i++ {
		e.intervals.Append(sorted[i].number, &sorted[i].frequency)

		if *debug {
			fmt.Fprintf(os.Stderr, "header: symbol=%d (%c), count=%d\n",
				sorted[i].number, sorted[i].number, sorted[i].count,
			)
		}

		_, err = output.Write([]byte{
			sorted[i].number,
			byte(0xFF & (sorted[i].count >> 0)),
			byte(0xFF & (sorted[i].count >> 8)),
			byte(0xFF & (sorted[i].count >> 16)),
			byte(0xFF & (sorted[i].count >> 32)),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeSymbols() error {
	for i := 0; i < len(e.data); i++ {
		var tsBegin, tsEnd int64
		if *timed {
			tsBegin = time.Now().UnixNano()
		}

		interval, err := e.intervals.Find(e.data[i])
		if err != nil {
			fmt.Println(err)
			return err
		}

		e.intervals.Resize(interval)

		if *timed {
			tsEnd = time.Now().UnixNano()
			fmt.Fprintf(os.Stderr, "%d,%d\n", i, tsEnd-tsBegin)
		}
	}

	if *debug {
		var start, end big.Rat
		e.intervals.GetStart(&start)
		e.intervals.GetEnd(&end)
		fmt.Fprintf(os.Stderr, "Final interval is [%s, %s]\n",
			start.FloatString(50),
			end.FloatString(50),
		)
	}

	return nil
}

func (e *Encoder) encodeBinary(output *bufio.Writer) error {
	base := big.NewInt(2)
	exponent := big.NewInt(0)
	increment := big.NewInt(1)
	power := big.NewInt(0)
	frac := big.NewRat(0, 1)
	result := big.NewRat(0, 1)

	bits := []bool{}
	for {
		exponent.Add(exponent, increment)
		power.Exp(base, exponent, nil)
		frac.SetFrac(increment, power)

		result.Add(result, frac)
		if e.intervals.LowerThan(result) {
			result.Sub(result, frac)
			bits = append(bits, false)
		} else {
			bits = append(bits, true)
		}

		if len(bits) == 8 {
			err := output.WriteByte(bitsToByte(&bits))
			if err != nil {
				return err
			}
		}

		if !e.intervals.GreaterThan(result) {
			break
		}
	}

	err := output.WriteByte(bitsToByte(&bits))
	if err != nil {
		return err
	}

	if *debug {
		fmt.Fprintf(os.Stderr, "Final number is %s\n", result.FloatString(100))
	}

	return nil
}

func encodeData(in *bufio.Reader, out *bufio.Writer) error {
	encoder := Encoder{}
	if err := encoder.getSymbols(in, out); err != nil {
		return err
	}

	if err := encoder.getIntervals(out); err != nil {
		return err
	}

	if err := encoder.encodeSymbols(); err != nil {
		return err
	}

	if err := encoder.encodeBinary(out); err != nil {
		return err
	}

	return nil
}
