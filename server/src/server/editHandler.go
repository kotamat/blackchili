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
	decoder := json.NewDecoder(r.Body)
	var body EditHandlerRequestStruct
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	LyricsInDb.Update(bson.M{"lyricsid": body.LyricsId}, bson.M{"$set": bson.M{"currentlyriclines": body.CurrentLyricLines}})

	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(EditHandlerResponseStruct{body.LyricsId})
}