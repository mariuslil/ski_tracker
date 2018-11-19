package main

import (
	"net/http"
	"time"
)

var dbTracks SnowTrackStorage
var dbSkiers SkiersStorage
var Skiers Skier
var SnowTracks SnowTrack
var startTime time.Time

func main() {
	startTime = time.Now()
	dbTracks = &SnowTracksMongoDB{"mongodb://olebgr:password1@ds261253.mlab.com:61253/snowtracks",
	"snowtracks", "tracks"}
	dbSkiers = &SkiersMongoDB{"mongodb://olebgr:password1@mongodb@ds211774.mlab.com:11774/skiers",
		"skiers", "skiers"}

	dbTracks.Init()
	http.HandleFunc("/snowtrack/track/", trackHandler)
	http.HandleFunc("/snowtrack/skier/", skierHandler)
	http.HandleFunc("/snowtracks/api", apiHandler)
	http.ListenAndServe(getPort(), nil)

}