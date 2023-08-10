package music

//  in this file we tried to read the logfile of netease to get now playing.

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getHistoryFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(homeDir, "AppData", "Local", "Netease", "CloudMusic", "webdata", "file", "history")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("cloudmusic data folder not found")
	}
	return path
}

func GetNowPlaying() (string, string) {
	path := getHistoryFile()
	trackInfo := make(map[string]interface{})
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 3200)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	readString := string(buffer)

	for i := 0; i < 4; i++ {
		var decodedJSON interface{}
		decoder := json.NewDecoder(strings.NewReader(readString[1:]))
		err := decoder.Decode(&decodedJSON)
		if err != nil {
			break
		}
		remainingBytes, _ := io.ReadAll(decoder.Buffered())
		readString = string(append([]byte{readString[0]}, remainingBytes...))
		for key, value := range decodedJSON.(map[string]interface{}) {
			trackInfo[key] = value
		}
	}

	if len(trackInfo) == 0 {
		return "", ""
	}

	trackName := trackInfo["track"].(map[string]interface{})["name"].(string)
	artistList := []string{}
	artists := trackInfo["track"].(map[string]interface{})["artists"].([]interface{})
	for _, artist := range artists {
		artistList = append(artistList, artist.(map[string]interface{})["name"].(string))
	}

	return trackName, strings.Join(artistList, "/ ")
}
