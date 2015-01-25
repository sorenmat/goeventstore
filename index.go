package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sync"
)

// IndexEntry in the index
type IndexEntry struct {
	EventNumber uint64
	Position    uint64
	Length      uint64
}

type Index struct {
	index []IndexEntry
	mu    sync.Mutex
	file  *os.File
}

func (index *Index) Init() {
	index.index = readIndexFromDisk()
	index.file = openFile("data/index.dat")
	fmt.Printf("%d\n", len(index.index))
}

func (index *Index) MaxEventNumber() int {
	return len(index.index)
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

func (index *Index) Get(number int) IndexEntry {
	return index.index[number]
}

func (index *Index) Append(entry IndexEntry) {
	index.mu.Lock()
	writeIndexEntry(entry, index.file)
	index.mu.Unlock()
	index.index = append(index.index, entry)
}

func (index *Index) Close() {
	index.file.Close()
}
func bytesToIndex(b []byte) IndexEntry {
	newEntry := IndexEntry{}
	newEntry.EventNumber = binary.LittleEndian.Uint64(b[0:8])
	newEntry.Position = binary.LittleEndian.Uint64(b[8:16])
	newEntry.Length = binary.LittleEndian.Uint64(b[16:24])
	return newEntry
}

func writeIndexEntry(index IndexEntry, indexFile *os.File) {
	indexFile.Write(indexToBytes(index))

}

func readIndexFromDisk() []IndexEntry {
	data, err := ioutil.ReadFile("data/index.dat")
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
	return result
}
