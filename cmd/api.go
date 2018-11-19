package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type apidata struct {
	uptime string `json:"uptime"`
	info string `json:"info"`
}

var api = apidata{
	uptime: timeFormat(startTime),
	info:"Service for tracks and skiers.",
}

func timeFormat (startTime time.Time) string{
	endTime := time.Now()
	return fmt.Sprintf("P%dY%dM%dDT%dH%dM%dS",
		endTime.Year()-startTime.Year(), endTime.Month()-startTime.Month(),
		endTime.Day()-startTime.Day(), endTime.Hour()-startTime.Hour(), endTime.Minute()-startTime.Minute(),
		endTime.Second()-startTime.Second())
}

func GetAPIinfo(w *http.ResponseWriter){
	api.uptime = timeFormat(startTime)
	json.NewEncoder(*w).Encode(api)
}

