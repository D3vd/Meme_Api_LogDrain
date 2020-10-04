package controllers

import (
	"net/http"

	"Meme_Api_LogDrain/parsers"
)

// ProcessLogsHandler - Process the logs data
func (c *Controllers) ProcessLogsHandler(w http.ResponseWriter, r *http.Request) {

	routerLogs := parsers.GetRouterLogs(r.Body)
	defer r.Body.Close()

	c.Mongo.UpdateDayAnalytics(routerLogs)
}
