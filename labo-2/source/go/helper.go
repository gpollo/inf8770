package main

import (
	"io"
	"os"
	"os/exec"
)

func ReadAll(r io.Reader) ([]byte, error) {
	data := []byte{}

	for {
		buffer := make([]byte, 1024)

		n, err := r.Read(buffer)
		if err != nil {
			return []byte{}, err
		}

		data = append(data, buffer[:n]...)
		if n != len(buffer) {
			break
		}
	}

	return data, nil
}

func CallProcess(args []string, data []byte) ([]byte, error) {
	inputR, inputW := io.Pipe()
	outputR, outputW := io.Pipe()

	command := exec.Command(args[0], args[1:]...)
	command.Stdin = inputR
	command.Stdout = outputW
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		panic(err.Error())
	}

	_, err := inputW.Write(data)
	if err != nil {
		return []byte{}, err
	}
	inputR.Close()
	inputW.Close()

	result, err := ReadAll(outputR)
	if err != nil {
		return []byte{}, err
	}
	outputR.Close()
	outputW.Close()

	command.Wait()
	return result, nil
}
