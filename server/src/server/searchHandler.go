package server

import (
	"fmt"
	"encoding/xml"
	"net/http"
	"net/url"
	"io/ioutil"
	"github.com/dustin/gojson"
	"gopkg.in/mgo.v2/bson"
)

type SearchRequestStruct struct {
	title string
}

type SearchResponseStruct struct {
	LyricsId string `json:"lyricsId"`
	LyricLines []LyricLine `json:"lines"`
}

type PetitLyrics struct {
	XMLName xml.Name `xml:"response"`
	PetitSongs PetitSongs `xml:"songs"`
}

type PetitSongs struct {
	XMLName xml.Name `xml:"songs"`
	PetitSongs []PetitSong `xml:"song"`
}

type PetitSong struct {
	XMLName xml.Name `xml:"song"`
	LyricsId string `xml:"lyricsId"`
	LyricsData string `xml:"lyricsData"`
}

type LyricsData struct {
	LyricLines []LyricLine `json:"lines"`
}

type LyricLine struct {
	Time int `json:"time"`
	Words []LyricWord `json:"words"`
}

type LyricWord struct {
	End int `json:"end"`
	Start int `json:"start"`
	Word string `json:"string"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	body := SearchRequestStruct{r.URL.Query().Get("title")}

	petitValues := url.Values{}
	petitValues.Add("title", body.title)
	petitValues.Add("priority", "3")
	petitValues.Add("auth_key", ServerConfig.PetitLyrics.AuthKey)

	res, err := http.Get("https://pl.t.petitlyrics.com/mh/1/lyrics/list.xml?" + petitValues.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	resBody, readErr := ioutil.ReadAll(res.Body);
	if readErr != nil {
		return
	}
	var petitLyrics PetitLyrics
	xml.Unmarshal(resBody, &petitLyrics)
	lyricsId := petitLyrics.PetitSongs.PetitSongs[0].LyricsId
	var lyricsTmp []SearchResponseStruct
	LyricsInDb.Find(bson.M{"lyricsid": lyricsId}).All(&lyricsTmp)

	var lyricsTimeData SearchResponseStruct
	if len(lyricsTmp) == 0 {
		var lyricsData LyricsData
		json.Unmarshal([]byte(petitLyrics.PetitSongs.PetitSongs[0].LyricsData), &lyricsData)
		lyricsTimeData = SearchResponseStruct{lyricsId, lyricsData.LyricLines}
		LyricsInDb.Insert(lyricsTimeData)
	} else {
		fmt.Println("data found in db")
		lyricsTimeData = lyricsTmp[0]
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(lyricsTimeData)

}
