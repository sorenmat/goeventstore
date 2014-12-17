package main

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestAuthorSerialization(t *testing.T) {
	author := AuthorType{Name: "Test Testsen"}
	data, _ := xml.Marshal(&author)
	xml := string(data)
	expectedXml := "<author><name>Test Testsen</name></author>"
	if xml != expectedXml {
		t.Error("Generated author xml is not correct, was ", xml)
	}
}

func TestEntrySerialization(t *testing.T) {
	author := AuthorType{Name: "Test Testsen"}

	dt := time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local)
	//fmt.Println(dt.Format(layout))
	entrytype := EntryType{Title: "Title", Id: "id", Updated: dt, Author: author, Summary: "summary", Link: "http://somewhere.com"}
	data, _ := xml.Marshal(&entrytype)
	xml := string(data)
	expectedXml := "<entry><title>Title</title><id>id</id><updated>2009-11-10T15:00:00+01:00</updated><author><name>Test Testsen</name></author><summary>summary</summary><link>http://somewhere.com</link></entry>"
	if xml != expectedXml {
		t.Errorf("Generated author xml is not correct, was \n%s \nshould have been \n%s", xml, expectedXml)
	}
}
