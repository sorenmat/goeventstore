package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"code.google.com/p/go-uuid/uuid"
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
*/
func TestWrite(t *testing.T) {
	LOOP_COUNT := 1000000
	eventStream := EventStream{}
	eventStream.Init()
	start := time.Now().UnixNano()
	eventData := "This is just a simple event, with some arbitary data in it...."
	myEvent := Event{Id: uuid.NewUUID(), EventType: "TestEvent", EventData: []byte(eventData)}
	for i := 0; i < LOOP_COUNT; i++ {
		eventStream.WriteEvent(myEvent)
	}
	fmt.Printf("%d writes took: %d sec\n", LOOP_COUNT, (time.Now().UnixNano()-start)/1000/1000)
}
func TestMultipleWrites(t *testing.T) {
	eventStream := EventStream{}
	eventStream.Init()
	fmt.Println("Starting multiple writes")
	eventData := "This is just a simple event, with some arbitary data in it...."
	myEvent := Event{Id: uuid.NewUUID(), EventType: "TestEvent", EventData: []byte(eventData)}

	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup
	go writestuff(1, eventStream, myEvent, &wg)
	go writestuff(2, eventStream, myEvent, &wg)
	go writestuff(3, eventStream, myEvent, &wg)
	go writestuff(4, eventStream, myEvent, &wg)
	wg.Wait()
}
func writestuff(c int, eventStream EventStream, myEvent Event, wg *sync.WaitGroup) {
	LOOP_COUNT := 10000
	wg.Add(1)
	for i := 0; i < LOOP_COUNT; i++ {
		eventStream.WriteEvent(myEvent)
		fmt.Printf("%d wrote %d\n", c, i)
		time.Sleep(2000)
	}
	wg.Done()
}

/*
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
	b.ResetTimer()
	eventData := string(time.Now().String())
	myEvent := Event{Id: uuid.NewUUID(), EventType: "TestEvent", EventData: []byte(eventData)}
	for i := 0; i < b.N; i++ {
		eventStream.WriteEvent(myEvent)
	}
	eventStream.Close()
}
