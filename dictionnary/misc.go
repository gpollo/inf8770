package main

import "io"
import "encoding/binary"

func WriteUvarint(w io.Writer, x uint64) error {
	buffer := make([]byte, 64)
	count := binary.PutUvarint(buffer, x)

	if _, err := w.Write(buffer[0:count]); err != nil {
		return err
	}

	return nil
}

func ReadUvarint(r io.ByteReader) (uint64, error) {
	return binary.ReadUvarint(r)
}
