package server

import (
	//"encoding/json"
	"net/http"
	"github.com/dustin/gojson"
	"gopkg.in/mgo.v2/bson"
)

type EditHandlerRequestStruct struct {
	LyricsId string `json:"lyricsId"`
	CurrentLyricLines []LyricLine `json:"currentLines"`
}

type EditHandlerResponseStruct struct {
	LyricsId string `json:"lyricsId"`
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var body EditHandlerRequestStruct
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	LyricsInDb.Update(bson.M{"lyricsid": body.LyricsId}, bson.M{"$set": bson.M{"currentlyriclines": body.CurrentLyricLines}})

	json.NewEncoder(w).Encode(EditHandlerResponseStruct{body.LyricsId})
}