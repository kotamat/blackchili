package main

import (
	"./server"
	"net/http"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"fmt"
	//"encoding/json"
)


func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aa")
}

func main() {
	server.LoadConfig()
	server.ConnectDatabase()

	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
