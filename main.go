package main

import (
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/log", processLogs)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
