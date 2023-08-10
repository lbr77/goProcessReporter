package api

import (
	"bytes"
	"fmt"
	"goProcessReporter/drivers/logger"
	"io"
	"net/http"
	"time"
)

// # in this file we try to send file to API configured as API_URL

func Report(processName, apiKey, apiUrl string, mediaTitle string, mediaArtist string) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0",
	}
	timestamp := int(time.Now().Unix())
	formedData := fmt.Sprintf(`{
		"timestamp": %d ,
		"process":   "%s",
		"key":       "%s"
	}`, timestamp, processName, apiKey)
	if mediaTitle != "" && mediaArtist != "" {
		formedData = fmt.Sprintf(`{
			"timestamp": %d ,
			"process":   "%s",
			"key":       "%s",
			"media": {
				"title": "%s",
				"artist": "%s"
			}
		}`, timestamp, processName, apiKey, mediaTitle, mediaArtist)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(formedData)))
	if err != nil {
		logger.Log.Error("Failed to create request,", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Error("failed to post", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("Failed to read response body,", err)
		return
	}
	logger.Log.Info(string(responseBody))
}
