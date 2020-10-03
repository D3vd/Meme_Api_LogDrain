package controllers

import (
	"log"
	"net/http"

	"Meme_Api_LogDrain/parsers"
)

// ProcessLogsHandler - Process the logs data
func ProcessLogsHandler(w http.ResponseWriter, r *http.Request) {

	logs := parsers.GetRouterLogs(r.Body)

	for _, line := range logs {
		log.Println(line)
	}
}
