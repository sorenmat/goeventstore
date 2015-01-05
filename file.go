package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"

	"code.google.com/p/go-uuid/uuid"

	"github.com/davecheney/profile"
	"github.com/golang/protobuf/proto"
)

// IndexEntry in the index
type IndexEntry struct {
	EventNumber int64
	Position    int64
	Length      int
}

// EventStream is the representation of the stream of events
type EventStream struct {
	index           []IndexEntry
	eFile           *os.File
	indexFile       *os.File
	indexEncoder    *gob.Encoder
	currentPosision int64
	maxEventNumber  int64
}

/**
Init needs to be called before using the eventstore
*/
func (e *EventStream) Init() {
	if e.eFile == nil {
		defer profile.Start(profile.CPUProfile).Stop()
		log.Printf("Opening data file\n")

		var err error
		e.eFile = openFile("data/events.dat")
		if err != nil {
			log.Fatal("Error creating file", err)
		}

		e.indexFile = openFile("data/index.dat")

		e.indexEncoder = gob.NewEncoder(e.indexFile)
		indexDecoder := gob.NewDecoder(e.indexFile)
		stillReading := true
		for stillReading {
			ie := IndexEntry{}
			decodeerr := indexDecoder.Decode(&ie)
			if decodeerr == io.EOF {
				stillReading = false
				log.Println("Done loading index")
			} else {
				e.index = append(e.index, ie)
				stillReading = true
			}
		}
		fmt.Printf("%d\n", len(e.index))
		e.maxEventNumber = int64(len(e.index))
		log.Printf("Starting event store with %d events\n", e.maxEventNumber)
		e.currentPosision, err = e.eFile.Seek(0, 2) // seek to the end of file
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (e *EventStream) WriteEvent(event Event) int64 {
	// proto stuff
	protoMessage := new(ProtoEvent)
	protoMessage.EventData = event.EventData
	protoMessage.Id = proto.String(event.Id.String())
	protoMessage.EventType = proto.String(event.EventType)
	protoBytes, _ := proto.Marshal(protoMessage)

	currentEventNumber := e.maxEventNumber
	e.maxEventNumber = e.maxEventNumber + 1

	// Write event to file
	length, err := e.eFile.WriteAt(protoBytes, e.currentPosision)
	e.index = append(e.index, IndexEntry{currentEventNumber, e.currentPosision, length})

	encodeerr := e.indexEncoder.Encode(IndexEntry{currentEventNumber, e.currentPosision, length})
	if encodeerr != nil {
		log.Fatal(err)
	}
	e.currentPosision = e.currentPosision + int64(length)
	if err != nil {
		log.Fatalln("encode:", err)
	}

	return currentEventNumber
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

	e.eFile.ReadAt(bytes, filePos)

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
