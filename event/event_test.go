package event

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/sorenmat/timer"

	"code.google.com/p/go-uuid/uuid"
)

func TestSizeOfUUIDInBytes(t *testing.T) {
	ud := uuid.NewUUID()

	res := make([]byte, 16)
	idBytes := []byte(ud)
	for k, v := range idBytes {
		res[k] = v
	}
	x := uuid.UUID(res)
	if !uuid.Equal(ud, x) {
		t.Error()
	}

	fmt.Println("x: ", x)

}

func TestSizeEventHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint32(10))
	fmt.Println("uint32 length: ", len(buf.Bytes()))
}

func TestSerializer(t *testing.T) {
	event := Event{Id: uuid.NewUUID(), EventType: "One", EventData: []byte("Hello world")}

	timer.Start()
	bytes := eventToBytes(event)
	fmt.Println("Serialized: ", timer.Stop())

	expectedBytes := (4 + 4) + len([]byte(event.Id)) + len("One") + len("Hello world")
	if len(bytes) != expectedBytes {
		t.Errorf("Wrong serialization size should be %d was %d,\n", expectedBytes, len(bytes))
	}

}

func TestSerializerBackAndForth(t *testing.T) {
	event := Event{Id: uuid.NewUUID(), EventType: "One", EventData: []byte("Hello world")}

	bytes := eventToBytes(event)

	newEvent := deserializeEvent(bytes)
	if !uuid.Equal(newEvent.Id, event.Id) {
		t.Errorf("The id's should match %v was %v", event.Id, newEvent.Id)
	}
	if newEvent.EventType != event.EventType {
		fmt.Println(len(newEvent.EventType))
		fmt.Println(len(event.EventType))
		t.Errorf("EventTypes doesn't match '%s' was '%s'", event.EventType, newEvent.EventType)
	}
	if string(newEvent.EventData) != string(event.EventData) {
		t.Error("Event data doesn't match  '%s' was '%s'", event.EventData, newEvent.EventData)
	}

}
