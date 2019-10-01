package main

import (
	"math/big"
	"sort"
)

type Decoder struct {
	total     int
	symbols   [SYMBOL_COUNT]Symbol
	intervals Intervals
	float     big.Rat
}

func (d *Decoder) ReadHeaderEntryCount(in Reader) (uint64, error) {
	count, err := ReadUvarint(in)
	if err != nil {
		return 0, err
	}

	debugHeaderEntryCount(count)

	return count, nil
}

func (d *Decoder) ReadHeaderEntry(in Reader) (byte, uint64, error) {
	symbol, err := in.ReadByte()
	if err != nil {
		return 0, 0, err
	}

	count, err := ReadUvarint(in)
	if err != nil {
		return 0, 0, err
	}

	debugHeaderEntry(symbol, count)

	return symbol, count, err
}

func (d *Decoder) GetSymbols(in Reader) error {
	d.total = 0
	for i := 0; i < SYMBOL_COUNT; i++ {
		d.symbols[i].number = 0
		d.symbols[i].count = 0
		d.symbols[i].frequency.SetInt64(0)
	}

	size, err := d.ReadHeaderEntryCount(in)
	if err != nil {
		return err
	}

	for i := 0; i < int(size); i++ {
		number, count, err := d.ReadHeaderEntry(in)
		if err != nil {
			return err
		}

		d.symbols[number].number = number
		d.symbols[number].count = int(count)
		d.total += int(count)
	}

	for i := 0; i < SYMBOL_COUNT; i++ {
		d.symbols[i].computeFrequency(d.total)
	}

	return nil
}

func (d *Decoder) GetFloat(in Reader) error {
	base := big.NewInt(2)
	exponent := big.NewInt(0)
	increment := big.NewInt(1)
	power := big.NewInt(0)
	frac := big.NewRat(0, 1)

	d.float.SetInt64(0)

	data := make([]byte, 1024)
	for {
		n, err := in.Read(data)
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
					d.float.Add(&d.float, frac)
				}
			}
		}
	}

	debugRatFloat(&d.float)

	return nil
}

func (d *Decoder) GetIntervals() {
	sorted := Frequencies{}
	for i := 0; i < SYMBOL_COUNT; i++ {
		if d.symbols[i].count == 0 {
			continue
		}

		sorted = append(sorted, d.symbols[i])
	}
	sort.Sort(sorted)

	for i := 0; i < len(sorted); i++ {
		d.intervals.Append(sorted[i].number, &sorted[i].frequency)
	}

	debugIntervals(d.intervals)
}

func (d *Decoder) Decode(out Writer) error {
	for i := 0; i < d.total; i++ {
		interval, err := d.intervals.FindByRat(&d.float)
		if err != nil {
			return err
		}

		d.intervals.Resize(interval)

		if _, err := out.Write([]byte{interval.number}); err != nil {
			return err
		}
	}

	return nil
}

func Decode(in Reader, out Writer) error {
	decoder := Decoder{}

	if err := decoder.GetSymbols(in); err != nil {
		return err
	}

	if err := decoder.GetFloat(in); err != nil {
		return err
	}

	decoder.GetIntervals()

	if err := decoder.Decode(out); err != nil {
		return err
	}

	return nil
}
