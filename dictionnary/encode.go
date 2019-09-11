package main

import (
	"io"
)

type Encoder struct {
	data     []byte
	searcher *Searcher
}

func NewEncoder() *Encoder {
	return &Encoder{
		data:     []byte{},
		searcher: NewSearcher(),
	}
}

func (e *Encoder) ReadAll(in io.Reader) error {
	buffer := make([]byte, 1024)
	for {
		n, err := in.Read(buffer)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}

		for i := 0; i < n; i++ {
			e.data = append(e.data, buffer[i])
		}
	}

	return nil
}

func (e *Encoder) WriteHeaderEntryCount(out io.Writer, count int) error {
	debugHeaderCount(uint64(count))
	if err := WriteUvarint(out, uint64(count)); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) WriteHeaderEntry(out io.Writer, symbol byte) error {
	debugHeaderEntry(symbol)
	if _, err := out.Write([]byte{symbol}); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) BuildInitialTable(out io.Writer) error {
	e.searcher.Reset()

	initials := []byte{}
	for _, symbol := range e.data {
		if e.searcher.Contains(symbol) {
			continue
		}

		initials = append(initials, symbol)
		if _, err := e.searcher.AddNode(symbol); err != nil {
			return err
		}

		e.searcher.Reset()
	}

	if err := e.WriteHeaderEntryCount(out, len(initials)); err != nil {
		return err
	}

	for _, symbol := range initials {
		if err := e.WriteHeaderEntry(out, symbol); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) WriteCurrentValue(out io.Writer) error {
	value, err := e.searcher.GetValue()
	if err != nil {
		return err
	}

	debugOutput(value)
	if err := WriteUvarint(out, value); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) ProcessData(out io.Writer) error {
	e.searcher.Reset()

	for _, symbol := range e.data {
		if e.searcher.Next(symbol) {
			continue
		}

		if err := e.WriteCurrentValue(out); err != nil {
			return err
		}

		if _, err := e.searcher.AddNode(symbol); err != nil {
			return err
		}

		e.searcher.Reset()
		e.searcher.Next(symbol)
	}

	if err := e.WriteCurrentValue(out); err != nil {
		return err
	}

	return nil
}

func Encode(in io.Reader, out io.Writer) error {
	encoder := NewEncoder()

	if err := encoder.ReadAll(in); err != nil {
		return err
	}

	if err := encoder.BuildInitialTable(out); err != nil {
		return err
	}

	if err := encoder.ProcessData(out); err != nil {
		return err
	}

	return nil
}
