package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprint(responseWriter, "Welcome!\n")
}

func Hello(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(responseWriter, "hello, %s!\n", params.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
