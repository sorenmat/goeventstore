package main

import "code.google.com/p/go-uuid/uuid"

type Event struct {
	Id        uuid.UUID
	EventType string
	EventData []byte
}

func eventToBytes(entry Event) []byte {

	length := 36 + 50 + len(entry.EventData)
	res := make([]byte, length)

	idStr := []byte(entry.Id) //.String()
	for k, v := range idStr {
		res[k] = v
	}

	// event type
	eventStr := entry.EventType
	len := len(eventStr)
	if len > 50 {
		len = 50
	}
	for i := 0; i < len; i++ {
		res[36+i] = eventStr[i]

	}

	//event data
	for k, v := range entry.EventData {
		res[86+k] = v
	}

	return res
}

/*
func bytesToIndex(b []byte) IndexEntry {
	newEntry := IndexEntry{}
	newEntry.EventNumber = binary.LittleEndian.Uint64(b[0:8])
	newEntry.Position = binary.LittleEndian.Uint64(b[8:16])
	newEntry.Length = binary.LittleEndian.Uint64(b[16:24])
	return newEntry
}
*/
