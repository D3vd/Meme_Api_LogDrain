package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"Meme_Api_LogDrain/controllers"
	"Meme_Api_LogDrain/db"

	"github.com/getsentry/sentry-go"
)

func main() {

	// Initialize MongoDB server
	m := &db.Mongo{}
	m.Init()

	// Initialize Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:   os.Getenv("SENTRY_DSN"),
		Debug: true,
	}); err != nil {
		log.Fatalln("Sentry Init: ", err)
	}

	defer sentry.Flush(2 * time.Second)

	c := controllers.Controllers{
		Mongo: m,
	}

	http.HandleFunc("/ping", c.PingHandler)
	http.HandleFunc("/log", c.ProcessLogsHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	log.Println("Log Drain Listening")
}
