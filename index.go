package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"os"
)

// IndexEntry in the index
type IndexEntry struct {
	EventNumber uint64
	Position    uint64
	Length      uint64
}

func indexToBytes(entry IndexEntry) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, entry.EventNumber)

	pos_b := make([]byte, 8)
	binary.LittleEndian.PutUint64(pos_b, entry.Position)

	len_b := make([]byte, 8)
	binary.LittleEndian.PutUint64(len_b, entry.Length)

	res := make([]byte, 24)
	for k, v := range b {
		res[k] = v
	}
	for k, v := range pos_b {
		res[8+k] = v
	}
	for k, v := range len_b {
		res[16+k] = v
	}
	return res
}

func bytesToIndex(b []byte) IndexEntry {
	newEntry := IndexEntry{}
	newEntry.EventNumber = binary.LittleEndian.Uint64(b[0:8])
	newEntry.Position = binary.LittleEndian.Uint64(b[8:16])
	newEntry.Length = binary.LittleEndian.Uint64(b[16:24])
	return newEntry
}

func writeIndexEntry(index IndexEntry, indexFile *os.File) {
	//var indexFile       *os.File
	//	indexFile := openFile("data/newindex.dat")
	indexFile.Write(indexToBytes(index))
	//	indexFile.Close()
}

func readIndexFromDisk() []IndexEntry {
	indexFile := openFile("data/newindex.dat")
	data, err := ioutil.ReadFile("data/newindex.dat")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loading index with length ", len(data))
	if math.Mod(float64(len(data)), 24) != 0 {
		log.Fatal("Index file seems to be corrupt !")
	}
	result := []IndexEntry{}
	for i := 0; i < len(data); i = i + 24 {
		slice := data[i : i+24]
		entry := bytesToIndex(slice)
		result = append(result, entry)
	}
	indexFile.Close()
	return result
}
