package main

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
)

type Encoder struct {
	data      []byte
	symbols   [SYMBOL_COUNT]Symbol
	intervals Intervals
}

func (e *Encoder) GetSymbols(in Reader) error {
	for {
		b, err := in.ReadByte()
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

	return nil
}

func (e *Encoder) WriteHeaderEntryCount(out Writer, count int) error {
	debugHeaderEntryCount(uint64(count))

	if err := WriteUvarint(out, uint64(count)); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) WriteHeaderEntry(out Writer, symbol byte, count uint64) error {
	debugHeaderEntry(symbol, count)

	if _, err := out.Write([]byte{symbol}); err != nil {
		return err
	}

	if err := WriteUvarint(out, uint64(count)); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) GetIntervals(out Writer) error {
	sorted := Frequencies{}
	for i := 0; i < len(e.symbols); i++ {
		if e.symbols[i].count == 0 {
			continue
		}

		sorted = append(sorted, e.symbols[i])
	}
	sort.Sort(sorted)

	if len(sorted) == 0 {
		return errors.New("Cannot have an empty symbol list")
	}

	if err := e.WriteHeaderEntryCount(out, len(sorted)); err != nil {
		return err
	}

	for i := 0; i < len(sorted); i++ {
		e.intervals.Append(sorted[i].number, &sorted[i].frequency)

		symbol := sorted[i].number
		count := uint64(sorted[i].count)
		if err := e.WriteHeaderEntry(out, symbol, count); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) EncodeSymbols() error {
	for i := 0; i < len(e.data); i++ {
		interval, err := e.intervals.Find(e.data[i])
		if err != nil {
			fmt.Println(err)
			return err
		}

		e.intervals.Resize(interval)
	}

	debugIntervals(e.intervals)

	return nil
}

func (e *Encoder) EncodeBinary(out Writer) error {
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
			if _, err := out.Write([]byte{bitsToByte(&bits)}); err != nil {
				return err
			}
		}

		if !e.intervals.GreaterThan(result) {
			break
		}
	}

	debugRatFloat(result)

	if _, err := out.Write([]byte{bitsToByte(&bits)}); err != nil {
		return err
	}

	return nil
}

func Encode(in Reader, out Writer) error {
	encoder := Encoder{}
	if err := encoder.GetSymbols(in); err != nil {
		return err
	}

	if err := encoder.GetIntervals(out); err != nil {
		return err
	}

	if err := encoder.EncodeSymbols(); err != nil {
		return err
	}

	if err := encoder.EncodeBinary(out); err != nil {
		return err
	}

	return nil
}
