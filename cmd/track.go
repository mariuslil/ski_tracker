package main

type SnowTracksMongoDB struct {
	DatabaseURL string
	DatabaseName string
	//TracksCollectionName string
}

type SnowTrackStorage interface {
	Init()
}

type SnowTrack struct {
	Name string `json:"name"`
	Mountain string `json:"mountain"`
	Length float32 `json: "length"`
}