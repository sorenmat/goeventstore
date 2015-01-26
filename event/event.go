package event

import (
	"encoding/binary"
	"fmt"

	"code.google.com/p/go-uuid/uuid"
)

type Event struct {
	Id        uuid.UUID
	EventType string
	EventData []byte
	/*
	   Length of EventType
	   Length of EventData
	*/
}

type EventHeader struct {
	typelength  uint32 // 4 bytes
	eventlength uint32 // 4 bytes
}

func getHeaderBytes(event Event) ([]byte, int, int) {
	eventTypeLength := len(event.EventType)
	eventDataLength := len(event.EventData)

	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, uint32(eventTypeLength))
	fmt.Println(b)
	binary.LittleEndian.PutUint32(b, uint32(eventDataLength))
	fmt.Println(b)

	return b, eventTypeLength, eventDataLength
}

func SerializeEvent(entry Event) []byte {
	eventTypeLength := len(entry.EventType)
	eventDataLength := len(entry.EventData)

	data_length := 24 + eventTypeLength + eventDataLength
	res := make([]byte, data_length)

	idBytes := []byte(entry.Id)

	byteCounter := 0

	// type length
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(eventTypeLength))
	for _, v := range b {
		res[byteCounter] = v
		byteCounter = byteCounter + 1
	}
	// data length
	binary.LittleEndian.PutUint32(b, uint32(eventDataLength))

	for _, v := range b {
		res[byteCounter] = v
		byteCounter = byteCounter + 1
	}
	for _, v := range idBytes {
		res[byteCounter] = v
		byteCounter = byteCounter + 1
	}

	// event type
	eventStr := entry.EventType
	for i := 0; i < eventTypeLength; i++ {
		res[byteCounter] = eventStr[i]
		byteCounter = byteCounter + 1
	}

	//event data
	for _, v := range entry.EventData {
		res[byteCounter] = v
		byteCounter = byteCounter + 1
	}

	return res
}

func DeserializeEvent(bytes []byte) Event {
	eventTypeLength := binary.LittleEndian.Uint32(bytes[:4])
	//	eventDataLength := binary.LittleEndian.Uint32(bytes[4:8])

	idBytes := bytes[8:24]
	typeEndPos := 24 + eventTypeLength
	eventTypeBytes := bytes[24:typeEndPos]
	eventDataBytes := bytes[typeEndPos:]

	event := Event{uuid.UUID(idBytes), string(eventTypeBytes), eventDataBytes}
	return event
}
