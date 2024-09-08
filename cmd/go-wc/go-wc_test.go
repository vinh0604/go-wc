package main

import (
	"strings"
	"testing"
)

func TestCountBytes(t *testing.T) {
	if countBytes(strings.NewReader("hello")) != 5 {
		t.Error("countBytes should return 5 for 'hello'")
	}

	if countBytes(strings.NewReader("hello\nworld")) != 11 {
		t.Error("countBytes should return 11 for 'hello\nworld'")
	}

	if countBytes(strings.NewReader("hêllo")) != 6 {
		t.Error("countBytes should return 6 for 'hêllo'")
	}
}

func TestCountLines(t *testing.T) {
	if countLines(strings.NewReader("hello")) != 1 {
		t.Error("countLines should return 1 for 'hello'")
	}

	if countLines(strings.NewReader("hello\nworld")) != 2 {
		t.Error("countLines should return 2 for 'hello\nworld'")
	}

	if countLines(strings.NewReader("hello\r\nworld")) != 2 {
		t.Error("countLines should return 2 for 'hello\r\nworld'")
	}
}

func TestCountWords(t *testing.T) {
	if countWords(strings.NewReader("hello")) != 1 {
		t.Error("countWords should return 1 for 'hello'")
	}

	if countWords(strings.NewReader("hello world")) != 2 {
		t.Error("countWords should return 2 for 'hello world'")
	}

	if countWords(strings.NewReader("hello  world")) != 2 {
		t.Error("countWords should return 2 for 'hello  world'")
	}

	if countWords(strings.NewReader("\thello  world")) != 2 {
		t.Error("countWords should return 2 for '\thello  world'")
	}

	if countWords(strings.NewReader("\thello \n world")) != 2 {
		t.Error("countWords should return 2 for '\thello \n world'")
	}

	if countWords(strings.NewReader("")) != 0 {
		t.Error("countWords should return 0 for ''")
	}

	if countWords(strings.NewReader("\t\n ")) != 0 {
		t.Error("countWords should return 0 for '\t\n '")
	}
}

func TestCountLocaleChars(t *testing.T) {
	if countLocaleChars(strings.NewReader("hello")) != 5 {
		t.Error("countLocaleChars should return 5 for 'hello'")
	}

	if countLocaleChars(strings.NewReader("hello\nworld")) != 11 {
		t.Error("countLocaleChars should return 11 for 'hello\nworld'")
	}

	if countLocaleChars(strings.NewReader("hêllo")) != 5 {
		t.Error("countLocaleChars should return 5 for 'hêllo'")
	}
}
