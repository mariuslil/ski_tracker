package main

import "time"

type SnowTracksMongoDB struct {
	DatabaseURL string
	DatabaseName string
	//TracksCollectionName string
}

type SnowTrackStorage interface {
	Init()
}

type SnowTrack struct {
	Mountain string `json:"Mountain"`
	Length float32 `json: "Length"`
	Time time.Duration `json: "Time"`
}