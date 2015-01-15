package main

import (
	"fmt"
	"testing"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

func TestSizeOfUUIDInBytes(t *testing.T) {
	ud := uuid.New()
	bytes := []byte(ud)
	fmt.Println(len(bytes))
}

func TestSerializer(t *testing.T) {
	event := Event{Id: uuid.NewUUID(), EventType: "One", EventData: []byte("Hello world")}

	start := time.Now().UnixNano()
	bytes := eventToBytes(event)
	fmt.Println("Serialized: ", (time.Now().UnixNano() - start))
	if len(bytes) != 97 {
		t.Error("Wrong serialization size should be 97 was,", len(bytes))
	}
}
