package db

import (
	"context"
	"log"
	"strings"
	"time"

	"Meme_Api_LogDrain/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateDayAnalytics : Updates values in Day Collection
func (m *Mongo) UpdateDayAnalytics(logs []types.RouterLog) {

	// Set Upsert = true so that it will create a new
	// document if the day entry is not present
	opts := options.Update().SetUpsert(true)

	// Get today's date and format
	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	currentDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	// Create Filter with todays date
	filter := bson.M{"date": currentDate}

	for _, line := range logs {
		update := bson.M{"$inc": bson.M{
			"statuses." + line.Status:                       1,
			"routes." + line.Path:                           1,
			"ip." + strings.Replace(line.Fwd, ".", "-", -1): 1}}

		_, err := m.DB.Collection("days").UpdateOne(context.TODO(), filter, update, opts)

		if err != nil {
			log.Println("Error: While updating analytics", err)
		}
	}
}
