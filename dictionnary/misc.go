package main

import "encoding/binary"

type Writer interface {
	Write(p []byte) (n int, err error)
}

func WriteUvarint(w Writer, x uint64) error {
	buffer := make([]byte, 64)
	count := binary.PutUvarint(buffer, x)

	if _, err := w.Write(buffer[0:count]); err != nil {
		return err
	}

	return nil
}

type Reader interface {
	Read(p []byte) (n int, err error)
	ReadByte() (byte, error)
}

func ReadUvarint(r Reader) (uint64, error) {
	return binary.ReadUvarint(r)
}
