package main

import (
	"bufio"
)

func Decode(in *bufio.Reader, out *bufio.Writer) error {
	count, err := ReadUvarint(in)
	if err != nil {
		return err
	}

	debugHeaderCount(count)

	searcher := NewSearcher()

	for i := 0; i < int(count); i++ {
		symbol, err := in.ReadByte()
		if err != nil {
			return err
		}

		debugHeaderEntry(symbol)
		if _, err = searcher.AddNode(symbol); err != nil {
			return err
		}
		searcher.Reset()
	}

	values, err := searcher.GetCurrentNode().GetValueToPathMap()
	if err != nil {
		return err
	}

	data := []byte{}
	for {
		var saved uint64
		for {
			value, err := ReadUvarint(in)
			if err != nil {
				return nil
			}
			debugOutput(value)

			if path, ok := values[value]; ok {
				data = append(data, path...)
				if _, err = out.Write(path); err != nil {
					return err
				}
			} else {
				saved = value
				break
			}
		}

		for i := 0; i < len(data); i++ {
			if !searcher.Next(data[i]) {
				value, err := searcher.AddNode(data[i])
				if err != nil {
					return err
				}
				debugNewEntry(searcher.GetCurrentPath(), value)

				if _, ok := values[value]; ok {
					panic("Malformed value dictionnary")
				}

				values[value] = searcher.GetCurrentPath()

				searcher.Reset()
				i--
			}
		}

		data = []byte{}

		if path, ok := values[saved]; ok {
			data = append(data, path...)
			if _, err = out.Write(path); err != nil {
				return err
			}
		} else {
			panic("Symbol not found after updating table")
		}

	}

	return nil
}
