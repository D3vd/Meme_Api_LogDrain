package main

import (
	"net/http"
	"os"

	"Meme_Api_LogDrain/controllers"
)

func main() {

	http.HandleFunc("/log", controllers.ProcessLogsHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
