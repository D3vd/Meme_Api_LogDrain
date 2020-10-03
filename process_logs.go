package main

import (
	"log"
	"net/http"
)

func processLogs(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
}
