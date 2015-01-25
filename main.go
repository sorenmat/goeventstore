package main

func main() {
	eventStream := EventStream{}
	eventStream.Init()

	server := RestServer{eventStream}
	server.Register()
	//	log.Fatal(http.ListenAndServe(":8080", nil))
}
