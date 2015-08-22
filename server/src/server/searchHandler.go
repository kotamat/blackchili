package server

import (
	"fmt"
	//"encoding/json"
	"net/http"
)

type SearchRequestStruct struct {
	title string
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	body := SearchRequestStruct{r.URL.Query().Get("title")}
	fmt.Println(body.title)
	// TODO call petitlyrics api
	// TODO call yahoo api
}
