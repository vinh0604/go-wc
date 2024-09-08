package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

type countFlags struct {
	bytes  bool
	lines  bool
	words  bool
	locale bool
}

type countBuffer struct {
	byteCount   int
	lineCount   int
	wordCount   int
	localeCount int
	lineBuf     []byte
	wordBuf     []byte
	localeBuf   []byte
}

func (cBuf *countBuffer) countLines(buf []byte, n int) {
	if n == 0 {
		return
	}

	lineBuf := append(cBuf.lineBuf, buf[:n]...)
	lines := bytes.Split(lineBuf, []byte{'\n'})
	cBuf.lineCount += len(lines) - 1
	cBuf.lineBuf = lines[len(lines)-1]
}

func (cBuf *countBuffer) countWords(buf []byte, n int) {
	wordBuf := append(cBuf.wordBuf, buf[:n]...)
	words := bytes.Fields(wordBuf)
	runes := bytes.Runes(wordBuf)
	if unicode.IsSpace(runes[len(runes)-1]) {
		cBuf.wordCount += len(words)
		cBuf.wordBuf = nil
	} else {
		cBuf.wordCount += len(words) - 1
		cBuf.wordBuf = words[len(words)-1]
	}
}

func (cBuf *countBuffer) countLocaleChars(buf []byte, n int) {
	localeBuf := append(cBuf.localeBuf, buf[:n]...)
	runes := bytes.Runes(localeBuf)
	cBuf.localeCount += len(runes) - 1
	cBuf.localeBuf = []byte(string(runes[len(runes)-1]))
}

func count(r io.Reader, cFlags countFlags) ([]int, error) {
	cBuf := countBuffer{}

	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			if cFlags.lines && len(cBuf.lineBuf) > 0 {
				cBuf.lineCount++
			}
			if cFlags.words && len(cBuf.wordBuf) > 0 {
				cBuf.wordCount++
			}
			if cFlags.locale && len(cBuf.localeBuf) > 0 {
				cBuf.localeCount++
			}
			break
		}
		if err != nil {
			return nil, err
		}

		if cFlags.bytes {
			cBuf.byteCount += n
		}

		if cFlags.lines {
			cBuf.countLines(buf, n)
		}

		if cFlags.words {
			cBuf.countWords(buf, n)
		}

		if cFlags.locale {
			cBuf.countLocaleChars(buf, n)
		}

	}

	result := make([]int, 0, 4)
	if cFlags.bytes {
		result = append(result, cBuf.byteCount)
	}
	if cFlags.lines {
		result = append(result, cBuf.lineCount)
	}
	if cFlags.words {
		result = append(result, cBuf.wordCount)
	}
	if cFlags.locale {
		result = append(result, cBuf.localeCount)
	}
	return result, nil
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
		result, err := count(os.Stdin, cFlags)
		if err != nil {
			println(fmt.Sprintf("go-wc: %s", err))
			exitCode = 1
		} else {
			printResult("", result)
		}
	} else {
		for _, arg := range args {
			result, err := countForFile(arg, cFlags)

			if err != nil {
				println(fmt.Sprintf("go-wc: %s: %s", arg, err))
				exitCode = 1
			} else {
				printResult(arg, result)
			}
		}
	}

	os.Exit(exitCode)
}

func printResult(arg string, result []int) {
	for i, cnt := range result {
		if i > 0 {
			print(" ")
		}
		print(cnt)
	}

	if arg != "" {
		print(" ")
		print(arg)
	}
	print("\n")
}

func countForFile(file string, flags countFlags) ([]int, error) {
	fileReader, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	return count(fileReader, flags)
}
