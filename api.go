package main

import (
	"fmt"
	"time"
)

type AuthorType struct {
	XMLName struct{} `xml:"author"`
	Name    string   `xml:"name"`
}

type EntryType struct {
	XMLName struct{}   `xml:"entry"`
	Title   string     `xml:"title"`
	Id      string     `xml:"id"`
	Updated time.Time  `xml:"updated"`
	Author  AuthorType `xml:"author"`
	Summary string     `xml:"summary"`
	Link    string     `xml:"link"`
}

type AtomFeed struct {
	XMLName struct{}   `xml:"atomfeed"`
	Title   string     `xml:"title"`
	Id      string     `xml:"id"`
	Updated time.Time  `xml:"updated"`
	Author  AuthorType `xml:"author"`
	Link    string     `xml:"link"`
	Entry   EntryType  `xml:"entry"`
}

func main() {
	author := AuthorType{Name: "Test Teseter"}
	fmt.Println(author)
}
