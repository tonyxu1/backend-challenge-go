package util

import (
	"bytes"
	"io"
	"math/big"
)

func BigIntFromInt(i int) *big.Int {
	return big.NewInt(int64(i))
}

func CountLines(r io.Reader) (int, error) {
	var count int
	var read int
	var err error
	var newline []byte = []byte("\n")

	buffer := make([]byte, 32*1024)

	for {
		read, err = r.Read(buffer)
		if err != nil {
			break
		}

		count += bytes.Count(buffer[:read], newline)
	}

	if err == io.EOF {
		return count, nil
	}

	return count, err
}
