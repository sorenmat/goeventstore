package main

import (
	"fmt"
	"log"
	"os"

	"code.google.com/p/go-uuid/uuid"

	"github.com/golang/protobuf/proto"
)

// EventStream is the representation of the stream of events
type EventStream struct {
	index           []IndexEntry
	eFile           *os.File
	indexFile       *os.File
	currentPosision uint64
	maxEventNumber  uint64
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

		e.indexFile = openFile("data/index.dat")

		e.index = readIndexFromDisk()
		fmt.Printf("%d\n", len(e.index))
		e.maxEventNumber = uint64(len(e.index))
		log.Printf("Starting event store with %d events\n", e.maxEventNumber)
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
	e.indexFile.Close()
	e.eFile.Close()

}

func (e *EventStream) WriteEvent(event Event) int64 {
	// proto stuff
	//start := time.Now().UnixNano()
	protoBytes := eventToBytes(event)
	//fmt.Println("Serializing event took: ", (time.Now().UnixNano() - start))
	currentEventNumber := e.maxEventNumber
	e.maxEventNumber = e.maxEventNumber + 1

	// Write event to file
	length, err := e.eFile.WriteAt(protoBytes, int64(e.currentPosision))
	e.index = append(e.index, IndexEntry{currentEventNumber, e.currentPosision, uint64(length)})
	writeIndexEntry(IndexEntry{currentEventNumber, e.currentPosision, uint64(length)}, e.indexFile)
	e.currentPosision = e.currentPosision + uint64(length)
	if err != nil {
		log.Fatalln("encode:", err)
	}

	return int64(currentEventNumber)
}

func (e *EventStream) AllEvents() []Event {
	result := []Event{}

	for i := 0; i < int(e.maxEventNumber); i++ {
		result = append(result, e.ReadEvent(1))
	}
	return result

}
func (e *EventStream) ReadEvent(number int) Event {
	entry := e.index[number]
	filePos := entry.Position

	event := Event{}
	bytes := make([]byte, entry.Length)

	e.eFile.ReadAt(bytes, int64(filePos))

	protoEvent := new(ProtoEvent)
	err := proto.Unmarshal(bytes, protoEvent)
	if err != nil {
		log.Println(entry.Position, entry.Length, number)
		log.Fatalf("ReadEvent Error: %s\nBytes: '%s'\n", err, bytes)
	}

	event.Id = uuid.Parse(protoEvent.GetId())
	event.EventType = protoEvent.GetEventType()
	event.EventData = protoEvent.GetEventData()
	//err := msgpack.Unmarshal(bytes, &event)

	return event
}

func openFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln("open:", err)
	}
	return file
}
