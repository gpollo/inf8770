package main

import (
	"io"
)

type Decoder struct {
	data     []byte
	values   map[uint64][]byte
	searcher *Searcher
}

func NewDecoder() *Decoder {
	return &Decoder{
		data:     []byte{},
		searcher: NewSearcher(),
	}
}

func (d *Decoder) ReadHeaderEntryCount(in Reader) (uint64, error) {
	count, err := ReadUvarint(in)
	if err != nil {
		return 0, err
	}

	debugHeaderCount(count)

	return count, nil
}

func (d *Decoder) ReadHeaderEntry(in Reader) (byte, error) {
	symbol, err := in.ReadByte()
	if err != nil {
		return 0, err
	}

	debugHeaderEntry(symbol)

	return symbol, nil
}

func (d *Decoder) BuildInitialTable(in Reader) error {
	count, err := d.ReadHeaderEntryCount(in)
	if err != nil {
		return err
	}

	for i := 0; i < int(count); i++ {
		symbol, err := d.ReadHeaderEntry(in)
		if err != nil {
			return err
		}

		if _, err = d.searcher.AddNode(symbol); err != nil {
			return err
		}

		d.searcher.Reset()
	}

	d.values, err = d.searcher.GetCurrentNode().GetValueToPathMap()
	if err != nil {
		return err
	}

	return nil
}

func (d *Decoder) ReadNextValue(in Reader) (uint64, error) {
	value, err := ReadUvarint(in)
	if err != nil {
		return 0, err
	}

	debugOutput(value)

	return value, nil
}

func (d *Decoder) DecodeValue(value uint64, out Writer) (bool, error) {
	if path, ok := d.values[value]; ok {
		if _, err := out.Write(path); err != nil {
			return false, err
		}

		d.data = append(d.data, path...)

		return true, nil
	}

	return false, nil
}

func (d *Decoder) DecodeUntilUnknown(in Reader, out Writer) (uint64, error) {
	for {
		value, err := d.ReadNextValue(in)
		if err != nil {
			return 0, err
		}

		ok, err := d.DecodeValue(value, out)
		if err != nil {
			return 0, err
		}

		if !ok {
			return value, nil
		}
	}
}

func (d *Decoder) UpdateTableFromDecodedData() error {
	for i := 0; i < len(d.data); i++ {
		if d.searcher.Next(d.data[i]) {
			continue
		}

		value, err := d.searcher.AddNode(d.data[i])
		if err != nil {
			return err
		}

		if _, ok := d.values[value]; ok {
			panic("Malformed value dictionnary")
		}

		d.values[value] = d.searcher.GetCurrentPath()

		d.searcher.Reset()
		i--
	}

	d.data = []byte{}

	return nil
}

func (d *Decoder) ProcessData(in Reader, out Writer) error {
	for {
		unknown, err := d.DecodeUntilUnknown(in, out)
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}

		if err = d.UpdateTableFromDecodedData(); err != nil {
			return err
		}

		ok, err := d.DecodeValue(unknown, out)
		if err != nil {
			return err
		}

		if !ok {
			panic("Symbol not found after updating table")
		}
	}
}

func Decode(in Reader, out Writer) error {
	decoder := NewDecoder()

	if err := decoder.BuildInitialTable(in); err != nil {
		return err
	}

	if err := decoder.ProcessData(in, out); err != nil {
		return err
	}

	return nil
}
