package main

import (
	"encoding/xml"
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
	XMLName struct{}    `xml:"feed"`
	Xmlns   string      `xml:"xmlns,attr"`
	Title   string      `xml:"title"`
	Id      string      `xml:"id"`
	Updated time.Time   `xml:"updated"`
	Author  AuthorType  `xml:"author"`
	Link    []Link      `xml:"link"`
	Entries []EntryType `xml:"entry"`
}

// Link is the representation of a http link in the atom feed
type Link struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr,omitempty"`
}
