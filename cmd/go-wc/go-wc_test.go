package main

import (
	"slices"
	"testing"
)

func TestCountForFile(t *testing.T) {
	res, err := countForFile("../../data/test.txt", countFlags{
		bytes:  true,
		lines:  true,
		words:  true,
		locale: true,
	})

	if err != nil {
		t.Error("countForFile should not return an error. Got:", err)
	}

	expected := []int{335042, 7145, 58164, 332146}
	if !slices.Equal(res, expected) {
		t.Errorf("countForFile should return %v. Got: %v", expected, res)
	}
}

func TestCountBufferCountLines(t *testing.T) {
	cBuf := countBuffer{}
	cBuf.countLines([]byte("hello\nwo"), 8)
	if cBuf.lineCount != 1 {
		t.Error("countLines should increase 1 for 'hello\\nwo', got:", cBuf.lineCount)
	}
	if !slices.Equal(cBuf.lineBuf, []byte("wo")) {
		t.Errorf("countLines should set lineBuf to 'wo', got: %v", cBuf.lineBuf)
	}

	cBuf.countLines([]byte("rld\n"), 4)
	if cBuf.lineCount != 2 {
		t.Error("countLines should increase 1 for 'rld\\n', got:", cBuf.lineCount)
	}
	if !slices.Equal(cBuf.lineBuf, []byte{}) {
		t.Errorf("countLines should set lineBuf to nil, got: %v", cBuf.lineBuf)
	}

	cBuf.countLines([]byte("\n"), 1)
	if cBuf.lineCount != 3 {
		t.Error("countLines should increase 1 for '\\n', got:", cBuf.lineCount)
	}
	if !slices.Equal(cBuf.lineBuf, []byte{}) {
		t.Errorf("countLines should set lineBuf to nil, got: %v", cBuf.lineBuf)
	}
}

func TestCountBufferCountWords(t *testing.T) {
	cBuf := countBuffer{}
	cBuf.countWords([]byte("hello"), 5)
	if cBuf.wordCount != 0 {
		t.Error("countWords should increase 0 for 'hello', got:", cBuf.wordCount)
	}
	if !slices.Equal(cBuf.wordBuf, []byte("hello")) {
		t.Errorf("countWords should set wordBuf to 'hello', got: %v", cBuf.wordBuf)
	}

	cBuf.countWords([]byte("  world"), 7)
	if cBuf.wordCount != 1 {
		t.Error("countWords should increase 1 for '  world', got:", cBuf.wordCount)
	}
	if !slices.Equal(cBuf.wordBuf, []byte("world")) {
		t.Errorf("countWords should set wordBuf to 'world', got: %v", cBuf.wordBuf)
	}

	cBuf.countWords([]byte("\t\n hello "), 9)
	if cBuf.wordCount != 3 {
		t.Error("countWords should increase 2 for '\\t\\n hello ', got:", cBuf.wordCount)
	}
	if !slices.Equal(cBuf.wordBuf, []byte{}) {
		t.Errorf("countWords should set wordBuf to nil, got: %v", cBuf.wordBuf)
	}

	cBuf.countWords([]byte("world \n\n"), 8)
	if cBuf.wordCount != 4 {
		t.Error("countWords should increase 1 for 'world \\n\\n', got:", cBuf.wordCount)
	}
	if !slices.Equal(cBuf.wordBuf, []byte{}) {
		t.Errorf("countWords should set wordBuf to nil, got: %v", cBuf.wordBuf)
	}

	cBuf.countWords([]byte("hel"), 3)
	if cBuf.wordCount != 4 {
		t.Error("countWords should increase 1 for 'hel', got:", cBuf.wordCount)
	}
	if !slices.Equal(cBuf.wordBuf, []byte("hel")) {
		t.Errorf("countWords should set wordBuf to 'hel', got: %v", cBuf.wordBuf)
	}

	cBuf.countWords([]byte("lo"), 2)
	if cBuf.wordCount != 4 {
		t.Error("countWords should increase 0 for 'lo', got:", cBuf.wordCount)
	}
	if !slices.Equal(cBuf.wordBuf, []byte("hello")) {
		t.Errorf("countWords should set wordBuf to 'hello', got: %v", cBuf.wordBuf)
	}
}

func TestCountBufferCountLocaleChars(t *testing.T) {
	cBuf := countBuffer{}
	cBuf.countLocaleChars([]byte("hello "), 6)
	if cBuf.localeCount != 5 {
		t.Error("countLocaleChars should increase 5 for 'hello ', got:", cBuf.localeCount)
	}
	if !slices.Equal(cBuf.localeBuf, []byte(" ")) {
		t.Errorf("countLocaleChars should set localeBuf to ' ', got: %v", cBuf.localeBuf)
	}

	cBuf.countLocaleChars([]byte("wôrld"), 6)
	if cBuf.localeCount != 10 {
		t.Error("countLocaleChars should increase 5 for 'wôrld', got:", cBuf.localeCount)
	}
	if !slices.Equal(cBuf.localeBuf, []byte{'d'}) {
		t.Errorf("countLocaleChars should set localeBuf to 'd', got: %v", cBuf.localeBuf)
	}
}
