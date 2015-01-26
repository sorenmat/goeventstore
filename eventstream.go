package main

import (
	"log"
	"os"
	"sync"

	"github.com/sorenmat/goeventstore/event"
)

// EventStream is the representation of the stream of events
type EventStream struct {
	mu              sync.Mutex
	eFile           *os.File
	index           Index
	currentPosision uint64
}

/**
Init needs to be called before using the eventstore
*/
func (e *EventStream) Init() {
	if e.eFile == nil {
		//		defer profile.Start(profile.CPUProfile).Stop()
		log.Printf("Opening data file\n")

		var err error
		e.eFile = openFile("data/events.dat")
		if err != nil {
			log.Fatal("Error creating file", err)
		}
		e.index = Index{}
		e.index.Init()
		log.Printf("Starting event store with %d events\n", e.index.MaxEventNumber())
		e.currentPosision = e.currentPosition() // seek to the end of file
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (e *EventStream) currentPosition() uint64 {
	currentPosision, _ := e.eFile.Seek(0, 2) // seek to the end of file
	return uint64(currentPosision)
}

func (e *EventStream) Close() {
	e.index.Close()
	e.eFile.Close()

}

func (e *EventStream) WriteEvent(evt event.Event) int64 {

	protoBytes := event.SerializeEvent(evt)

	currentEventNumber := e.index.MaxEventNumber()

	// Write event to file
	e.mu.Lock()
	length, err := e.eFile.WriteAt(protoBytes, int64(e.currentPosision))
	e.mu.Unlock()
	len := uint64(length)

	e.index.Append(IndexEntry{uint64(currentEventNumber), e.currentPosision, len})

	e.currentPosision = e.currentPosision + len
	if err != nil {
		log.Fatalln("encode:", err)
	}

	return int64(currentEventNumber)
}

func (e *EventStream) AllEvents() []event.Event {
	result := []event.Event{}

	for i := 0; i < e.index.MaxEventNumber(); i++ {
		result = append(result, e.ReadEvent(1))
	}
	return result

}

func (e *EventStream) ReadEvent(number int) event.Event {
	entry := e.index.Get(number)
	filePos := entry.Position

	bytes := make([]byte, entry.Length)

	e.eFile.ReadAt(bytes, int64(filePos))

	evt := event.DeserializeEvent(bytes)

	return evt
}

func openFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln("open:", err)
	}
	return file
}
