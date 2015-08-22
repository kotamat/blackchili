package server
import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	PetitLyrics PetitLyricsConfig `json:"petitLyrics"`
}

type DatabaseConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

type PetitLyricsConfig struct {
	AuthKey string `json:"authKey"`
}

var ServerConfig Config

func LoadConfig() {
	absPath, _ := filepath.Abs("server/src/resources/params.json")
	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &ServerConfig)
}