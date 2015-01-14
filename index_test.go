package main

import (
	"fmt"
	"testing"
)

func TestIndexEntryToBytes(t *testing.T) {
	entry := IndexEntry{EventNumber: 1, Position: 0, Length: 20}
	bytes := indexToBytes(entry)
	if len(bytes) != 24 {
		t.Errorf("Wrong size of byte array was %d, should have been %d\n", len(bytes), 128)
	}

	entry = IndexEntry{EventNumber: 1202020, Position: 20202020, Length: 200000}
	bytes = indexToBytes(entry)
	if len(bytes) != 24 {
		t.Errorf("Wrong size of byte array was %d, should have been %d\n", len(bytes), 128)
	}

}

func TestWriteReadIndex(t *testing.T) {
	entry := IndexEntry{EventNumber: 1, Position: 0, Length: 20}
	bytes := indexToBytes(entry)
	if len(bytes) != 24 {
		t.Errorf("Wrong size of byte array was %d, should have been %d\n", len(bytes), 128)
	}
	index := bytesToIndex(bytes)
	fmt.Println(index)
	if index.EventNumber != 1 {
		t.Error("EventNumber should be 1")
	}
	if index.Length != 20 {
		t.Error("EventNumber should be 20")
	}
}

func BenchmarkWriteIndexEntries(b *testing.B) {
	entry := IndexEntry{EventNumber: 1, Position: 0, Length: 20}
	for i := 0; i < b.N; i++ {
		writeIndexEntry(entry)
	}
}
