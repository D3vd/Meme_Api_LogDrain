package types

import "time"

// Day based Analytics
type Day struct {
	Date     time.Time      `bson:"date"`
	IP       map[string]int `bson:"ip"`
	Statuses map[string]int `bson:"statuses"`
	Routes   map[string]int `bson:"routes"`
}
