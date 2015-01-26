package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sorenmat/goeventstore/event"
)

// RestServer struct to define the server
type RestServer struct {
	Server EventStream
}

// AllEvents return a atom feed of all the events
func (u RestServer) AllEvents(c *gin.Context) {
	url := "http://localhost:8080/event"
	title := "EventStream"
	updated := time.Now() // should be latest entry in stream
	author := AuthorType{Name: "GoEventStore"}

	link := Link{Href: "http://localhost:8080/event", Rel: "self"}
	links := []Link{link}
	//	u.Server.Init()
	events := u.Server.AllEvents()
	entries := eventsToAtom(events)
	//	fmt.Println(entries)
	feed := AtomFeed{Xmlns: "http://www.w3.org/2005/Atom", Title: title, Id: url, Updated: updated, Author: author, Link: links, Entries: entries}
	//response.WriteEntity(feed)
	c.JSON(200, feed)
}

func eventsToAtom(events []event.Event) []EntryType {
	result := []EntryType{}
	for _, v := range events {
		e := EntryType{Id: v.Id.String(), Link: "http://localhost:8080/event" + v.Id.String()}
		result = append(result, e)
	}
	return result
}

// AllEvents return a atom feed of all the events
func (u RestServer) addEventToStream(c *gin.Context) {
	var entity = event.Event{}
	if c.EnsureBody(&entity) {
		u.Server.WriteEvent(entity)

	}
	c.JSON(200, nil)
}

func (u RestServer) getEvent(c *gin.Context) {
	ids := c.Params.ByName("id")
	eventNumber, _ := strconv.Atoi(ids)

	c.JSON(200, u.Server.ReadEvent(eventNumber))
}

func (u RestServer) status(c *gin.Context) {
	c.JSON(200, "ok")

}

// Register all rest resources on the server
func (u RestServer) Register() {

	ws := gin.Default()
	ws.GET("/event", u.AllEvents)

	ws.GET("/status", u.status)
	ws.GET("/stream/:stream/:id", u.getEvent)
	ws.POST("/stream/:streamid", u.addEventToStream)

	ws.Run(":8080")
}
