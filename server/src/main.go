package main

import (
	"./server"
	"net/http"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"encoding/json"
)

func main() {
	server.LoadConfig()
	server.ConnectDatabase()

	http.HandleFunc("/search", server.SearchHandler)
	http.HandleFunc("/edit", server.EditHandler)
	http.ListenAndServe(":8080", nil)
}
