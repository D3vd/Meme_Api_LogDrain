package controllers

import (
	"log"
	"net/http"
)

// PingHandler Handles /ping Requests
func (c *Controllers) PingHandler(w http.ResponseWriter, r *http.Request) {
	IPAddress := readUserIP(r)

	log.Println("Ping Received From " + IPAddress)
	w.Write([]byte("OK"))
}

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
