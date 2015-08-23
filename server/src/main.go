package main

import (
	"./server"
	"net/http"
	"gopkg.in/mgo.v2"
)



func main() {
	server.LoadConfig()
	server.ConnectDatabase()
	server.LyricsInDb = server.Database.C(server.LYRICS_COLLECTION)
	lyricsIdIndex := mgo.Index{
		Key:	[]string{"lyricsid"},
		Unique:	true,
		DropDups:	true,
		Background:	true,
		Sparse:	true,
	}

	for _, v := range []mgo.Index{lyricsIdIndex} {
		_ = server.LyricsInDb.EnsureIndex(v)
	}

	http.HandleFunc("/search", server.SearchHandler)
	http.HandleFunc("/edit", server.EditHandler)
	http.ListenAndServe(":8080", nil)
}
