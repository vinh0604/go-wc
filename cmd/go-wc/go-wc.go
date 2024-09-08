package main

import (
	"bufio"
	"io"
)

func main() {

}

func countBytes(r io.Reader) int {
	buf := make([]byte, 4096)
	total := 0
	for {
		n, err := r.Read(buf)
		total += n
		if err == io.EOF {
			break
		}
	}

	return total
}

func countLines(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	total := 0

	for scanner.Scan() {
		total++
	}

	return total
}

func countWords(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	total := 0

	for scanner.Scan() {
		total++
	}

	return total
}

func countLocaleChars(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	total := 0

	for scanner.Scan() {
		total++
	}

	return total
}
