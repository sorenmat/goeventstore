package main

import (
	"log"
	"net/http"
)

func main() {
	eventStream := EventStream{}
	eventStream.Init()

	server := RestServer{}
	server.Server = eventStream
	server.Register()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
