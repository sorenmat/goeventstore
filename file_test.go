package main

import (
	"testing"
	"time"
)

/*
func TestWriteOneReadOne(t *testing.T) {
	eventStream := EventStream{}
	eventStream.Init()

	eventData := "This is just a simple event, with some arbitary data in it...."
	myEvent := Event{EventType: "TestEvent", EventData: []byte(eventData)}

	eventStream.WriteEvent(myEvent)
	event := eventStream.ReadEvent(0)
	if event.EventType != "TestEvent" {
		t.Error("Death..")
	}
}

func TestWrite(t *testing.T) {
	eventStream := EventStream{}
	eventStream.Init()
	fmt.Println("This is a test")
	eventData := "This is just a simple event, with some arbitary data in it...."
	myEvent := Event{EventType: "TestEvent", EventData: []byte(eventData)}
	for i := 0; i < 5000; i++ {
		eventStream.WriteEvent(myEvent)
	}
}

func TestAllEvents(t *testing.T) {
	eventStream := EventStream{}
	eventStream.Init()
	events := eventStream.AllEvents()
	if len(events) <= 0 {
		t.Error("Stream is empty")
	}
}
func TestWriteManyEvents(t *testing.T) {
	eventStream := EventStream{}
	eventStream.Init()
	start := time.Now()
	reportUnit := float64(100000)
	for i := 0; i < 1000*1000; i++ {
		if math.Mod(float64(i), reportUnit) == 0 {
			fmt.Printf("Saved: %.0f, and it took %d ns.\n", reportUnit, time.Now().Sub(start).Nanoseconds())
			start = time.Now()

		}
		event := Event{EventType: "TestEvent", EventData: []byte("This is a simple event")}
		eventStream.WriteEvent(event)
	}
}
func TestWriteAndRead(t *testing.T) {
	eventStream := EventStream{}
	eventStream.Init()
	fmt.Println("This is a test")
	for i := 0; i < 5000; i++ {
		eventData := "This is just a simple event, with some arbitary data in it...." + string(i)
		myEvent := Event{EventType: "TestEvent", EventData: []byte(eventData)}

		eventStream.WriteEvent(myEvent)
		eventStream.ReadEvent(i)
	}
}
*/

func BenchmarkWriteEvents(b *testing.B) {
	eventStream := EventStream{}
	eventStream.Init()
	b.StartTimer()
	eventData := string(time.Now().String())
	myEvent := Event{EventType: "TestEvent", EventData: []byte(eventData)}
	for i := 0; i < b.N; i++ {
		eventStream.WriteEvent(myEvent)
	}
	b.StopTimer()
	eventStream.Close()
}
