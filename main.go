package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type limitError struct {
	message    string
	limit      int
	lastString int
}

func (e *limitError) Error() string {
	return fmt.Sprintf("%s, limit: %d, last string: %v", e.message, e.limit, e.lastString)
}

func main() {
	limit := 2

	count, err := readerCountLines(limit)
	if err != nil {
		var enf *limitError
		if errors.As(err, &enf) {
			fmt.Println("string count exceed limit, please read another file =) err: ", err.Error())
			return
		}

		log.Fatal(err)
	}

	fmt.Printf("Total strings: %d\n", count)
}

func readerCountLines(limit int) (count int, err error) {
	count = 0

	in, err := os.Open("in.txt")
	if err != nil {
		return count, fmt.Errorf("can't open file: %s", err)
	}
	defer in.Close()

	reader := bufio.NewReader(in)

	for {
		count++
		if count > limit {
			err = &limitError{message: "over the limit", limit: limit, lastString: count - 1}
			break
		}
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
