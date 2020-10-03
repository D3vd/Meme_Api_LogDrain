package main

import (
	"log"
	"net/http"
	"os"

	"Meme_Api_LogDrain/controllers"
	"Meme_Api_LogDrain/db"
)

func main() {

	// Initialize MongoDB server
	m := &db.Mongo{}
	m.Init()

	c := controllers.Controllers{
		Mongo: m,
	}

	http.HandleFunc("/log", c.ProcessLogsHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

	log.Println("Log Drain Listening")
}
