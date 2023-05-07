package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	count, err := readerCountLines()
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Printf("Total strings: %d", count)
}

func readerCountLines() (count int, err error) {
	count = 0

	in, err := os.Open("in.txt")
	if err != nil {
		return count, fmt.Errorf("can't open file: %s", err)
	}
	defer in.Close()

	reader := bufio.NewReader(in)

	for {
		count++
		_, err = reader.ReadString(byte('\n'))
		if err != io.EOF && err != nil {
			err = fmt.Errorf("cant'read line %v: %s", count, err)
			break
		}
		if err == io.EOF {
			err = nil
			break
		}
	}
	return count, err
}
