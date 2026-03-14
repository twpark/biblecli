package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Reference struct {
	Book       *Book
	Chapter    int
	VerseStart int // 0 = whole chapter
	VerseEnd   int
}

var (
	reKorean  = regexp.MustCompile(`^([가-힣]+)(\d+.*)$`)
	reEnglish = regexp.MustCompile(`^(\d*[a-zA-Z]+)(\d+.*)$`)
	reChVerse = regexp.MustCompile(`^(\d+)(?::(\d+)(?:-(\d+))?)?$`)
)

func parseReference(input string) (*Reference, error) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	var bookName, rest string

	if m := reKorean.FindStringSubmatch(input); m != nil {
		bookName, rest = m[1], m[2]
	} else if m := reEnglish.FindStringSubmatch(input); m != nil {
		bookName, rest = m[1], m[2]
	} else {
		return nil, fmt.Errorf("invalid reference: %s", input)
	}

	book := lookupBook(bookName)
	if book == nil {
		return nil, fmt.Errorf("unknown book: %s", bookName)
	}

	m := reChVerse.FindStringSubmatch(rest)
	if m == nil {
		return nil, fmt.Errorf("invalid chapter/verse format: %s", rest)
	}

	chapter, _ := strconv.Atoi(m[1])
	if chapter == 0 {
		return nil, fmt.Errorf("invalid chapter: 0")
	}

	var verseStart, verseEnd int
	if m[2] != "" {
		verseStart, _ = strconv.Atoi(m[2])
		verseEnd = verseStart
	}
	if m[3] != "" {
		verseEnd, _ = strconv.Atoi(m[3])
	}

	if verseEnd < verseStart {
		verseStart, verseEnd = verseEnd, verseStart
	}

	return &Reference{
		Book:       book,
		Chapter:    chapter,
		VerseStart: verseStart,
		VerseEnd:   verseEnd,
	}, nil
}
