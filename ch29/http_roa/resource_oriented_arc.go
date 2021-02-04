package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employeeDB map[string]*Employee

func init() {
	employeeDB = map[string]*Employee{}
	employeeDB["Mike"] = &Employee{"e-1", "Mike", 35}
	employeeDB["Rose"] = &Employee{"e-2", "Rose", 45}
}

func Index(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprint(responseWriter, "Welcome!\n")
}

func GetEmployeeByName(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	queryName := params.ByName("name")
	var (
		ok       bool
		info     *Employee
		infoJson []byte
		err      error
	)

	if info, ok = employeeDB[queryName]; !ok {
		responseWriter.Write([]byte("{\"error\":\"Not Found\"}"))
		return
	}

	if infoJson, err = json.Marshal(info); err != nil {
		responseWriter.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err)))
		return
	}
	responseWriter.Write(infoJson)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/employees/:name", GetEmployeeByName)

	log.Fatal(http.ListenAndServe(":8080", router))

}
