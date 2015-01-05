package main

import (
	"time"

	"github.com/emicklei/go-restful"
)

// RestServer struct to define the server
type RestServer struct {
	Server EventStream
}

// AllEvents return a atom feed of all the events
func (u RestServer) AllEvents(request *restful.Request, response *restful.Response) {
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
	response.WriteEntity(feed)
}

func eventsToAtom(events []Event) []EntryType {
	result := []EntryType{}
	for _, v := range events {
		e := EntryType{Id: v.Id.String(), Link: "http://localhost:8080/event" + v.Id.String()}
		result = append(result, e)
	}
	return result
}

// Register all rest resources on the server
func (u RestServer) Register() {

	ws := new(restful.WebService)
	ws.
		Path("/").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("/event").To(u.AllEvents).
		// docs
		Doc("List all events").
		Operation("AllEvents").
		Writes(AtomFeed{})) // on the response

	/*ws.Route(ws.GET("/").To(u.findAllUsers).
	// docs
	Doc("get all users").
	Operation("findAllUsers").
	Returns(200, "OK", []User{}))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		// docs
		Doc("get a user").
		Operation("findUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(User{})) // on the response
	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
		// docs
		Doc("update a user").
		Operation("updateUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Reads(User{})) // from the request
	ws.Route(ws.POST("/{id}").To(u.createProperty).
		// docs
		Doc("create a property").
		Operation("createProperty").
		Reads(User{})) // from the request
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
		// docs
		Doc("delete a user").
		Operation("removeUser").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))
	*/
	restful.Add(ws)
}
