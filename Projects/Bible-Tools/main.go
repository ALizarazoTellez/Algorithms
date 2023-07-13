package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"

	_ "embed"
)

func main() {
	var bible Bible

	r := bytes.NewReader(bibleXML)
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel

	if err := dec.Decode(&bible); err != nil {
		panic(err)
	}

	fmt.Println(bible.getVerses(parseCite((os.Args[1]))))
}

//go:embed data/rv1960.xml
var bibleXML []byte

type Bible struct {
	Books []Book `xml:"b"`
}

func (bib *Bible) getVerses(book string, chapter, verseStart, verseEnd int) string {
	b := bib.Book(book)
	if b == nil {
		return ""
	}

	c := b.Chapter(chapter)
	if c == nil {
		return ""
	}

	var verses strings.Builder

	for i := verseStart; i <= verseEnd; i++ {
		verses.WriteString(c.Verse(i))
	}

	return verses.String()
}

func (b *Bible) Book(s string) *Book {
	for _, b := range b.Books {
		if b.Name == s {
			return &b
		}
	}

	return nil
}

type Book struct {
	Name     string    `xml:"n,attr"`
	Chapters []Chapter `xml:"c"`
}

func (b *Book) Chapter(n int) *Chapter {
	for _, c := range b.Chapters {
		if c.Number == n {
			return &c
		}
	}

	return nil
}

type Chapter struct {
	Number int     `xml:"n,attr"`
	Verses []Verse `xml:"v"`
}

func (c *Chapter) Verse(n int) string {
	for _, v := range c.Verses {
		if v.Number == n {
			return v.Text
		}
	}

	return ""
}

type Verse struct {
	Number int    `xml:"n,attr"`
	Text   string `xml:",chardata"`
}

// Thanks ChatGPT!.
var reCite = regexp.MustCompile(`\b(\d?\s?[a-zA-ZáéíóúÁÉÍÓÚüÜ]+\s?[a-zA-ZáéíóúÁÉÍÓÚüÜ]+(?:\s?[a-zA-ZáéíóúÁÉÍÓÚüÜ]+)?)(?:\s+(\d+)(?::(\d+))?(?:-(\d+))?)?\b`)

func parseCite(cite string) (book string, chapter, verseStart, verseEnd int) {
	parts := reCite.FindStringSubmatch(cite)[1:]

	book = parts[0]

	chapter, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	verseStart, err = strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	verseEnd, err = strconv.Atoi(parts[3])
	if err != nil {
		verseEnd = verseStart
	}

	return book, chapter, verseStart, verseEnd
}
