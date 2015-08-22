package server

import (
	"fmt"
	"encoding/xml"
	"net/http"
	"net/url"
	"io/ioutil"
	"github.com/dustin/gojson"
)

type SearchRequestStruct struct {
	title string
}

type SearchResponseStruct struct {
	TimeData LyricsData `json:"timeData`
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
	resBody, readerr := ioutil.ReadAll(res.Body);
	if readerr != nil {
		return
	}
	var lyrics PetitLyrics
	xml.Unmarshal(resBody, &lyrics)
	var lyricsTimeData LyricsData
	fmt.Println(lyrics.PetitSongs.PetitSongs[0].LyricsData)
	json.Unmarshal([]byte(lyrics.PetitSongs.PetitSongs[0].LyricsData), &lyricsTimeData)

	// TODO call yahoo api

	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := SearchResponseStruct{lyricsTimeData}
	json.NewEncoder(w).Encode(result)

}
