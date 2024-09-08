package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

type countFlags struct {
	bytes  bool
	lines  bool
	words  bool
	locale bool
}

func main() {
	bytes := flag.Bool("c", false, "count bytes")
	lines := flag.Bool("l", false, "count lines")
	words := flag.Bool("w", false, "count words")
	locale := flag.Bool("m", false, "count locale chars")

	flag.Parse()

	var cFlags countFlags
	if !*bytes && !*lines && !*words && !*locale {
		cFlags = countFlags{
			bytes:  true,
			lines:  true,
			words:  true,
			locale: false,
		}
	} else {
		cFlags = countFlags{
			bytes:  *bytes,
			lines:  *lines,
			words:  *words,
			locale: *locale,
		}
	}

	args := flag.Args()

	exitCode := 0
	if len(args) == 0 {
		result, err := countForStdIn(cFlags)
		if err != nil {
			println(fmt.Sprintf("go-wc: %s", err))
			exitCode = 1
		} else {
			printResult("", result, false)
		}
	} else {
		showFilename := len(args) > 1
		for _, arg := range args {
			result, err := countForFile(arg, cFlags)

			if err != nil {
				println(fmt.Sprintf("go-wc: %s: %s", arg, err))
				exitCode = 1
			} else {
				printResult(arg, result, showFilename)
			}
		}
	}

	os.Exit(exitCode)
}

func printResult(arg string, result []int, showFilename bool) {
	if len(result) > 1 {
		for _, cnt := range result {
			print(fmt.Sprintf("\t%d", cnt))
		}
	} else if len(result) == 1 {
		print(result[0])
	}

	if showFilename {
		print(fmt.Sprintf("\t%s", arg))
	}

	print("\n")
}

func countForFile(file string, flags countFlags) ([]int, error) {
	fileReader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	result := make([]int, 0, 4)

	if flags.bytes {
		fileReader.Seek(0, 0)
		result = append(result, countBytes(fileReader))
	}

	if flags.lines {
		fileReader.Seek(0, 0)
		result = append(result, countLines(fileReader))
	}

	if flags.words {
		fileReader.Seek(0, 0)
		result = append(result, countWords(fileReader))
	}

	if flags.locale {
		fileReader.Seek(0, 0)
		result = append(result, countLocaleChars(fileReader))
	}

	return result, nil
}

func countForStdIn(flags countFlags) ([]int, error) {
	var buffer bytes.Buffer
	// Copy from stdin to buffer
	// This is done to be able to read the StdIn multiple times
	// TODO: Implement multple count flags from a single read
	_, err := io.Copy(&buffer, os.Stdin)
	if err != nil {
		return nil, err
	}

	result := make([]int, 0, 4)

	if flags.bytes {
		result = append(result, countBytes(bytes.NewReader(buffer.Bytes())))
	}

	if flags.lines {
		result = append(result, countLines(bytes.NewReader(buffer.Bytes())))
	}

	if flags.words {
		result = append(result, countWords(bytes.NewReader(buffer.Bytes())))
	}

	if flags.locale {
		result = append(result, countLocaleChars(bytes.NewReader(buffer.Bytes())))
	}

	return result, nil
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
