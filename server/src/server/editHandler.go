package server

import (
	"fmt"
	//"encoding/json"
	"net/http"
	"github.com/dustin/gojson"
)

type EditHandlerRequestStruct struct {
	Abc string
}

type EditHandlerResponseStruct struct {

}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body EditHandlerRequestStruct
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body.Abc)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// TODO
}