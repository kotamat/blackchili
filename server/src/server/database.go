package server

import (
	"gopkg.in/mgo.v2"
)

var DatabaseSession *mgo.Session
var Database *mgo.Database
var LyricsInDb *mgo.Collection
const LYRICS_COLLECTION string = "lyrics"

func ConnectDatabase() {
	session, err := mgo.Dial(ServerConfig.Database.Host)
	DatabaseSession = session
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	Database = session.DB(ServerConfig.Database.DatabaseName)
}