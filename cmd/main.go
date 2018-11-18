package main

import (
	"net/http"
)

var dbTracks SnowTrackStorage
var dbSkiers SkiersStorage
var Skiers Skier
var SnowTracks SnowTrack

func main() {
	dbTracks = &SnowTracksMongoDB{"mongodb://olebgr:password1@ds261253.mlab.com:61253/snowtracks",
	"snowtracks", "tracks"}
	dbSkiers = &SkiersMongoDB{"mongodb://olebgr:password1@ds261253.mlab.com:61253/snowtracks",
		"skiers", "skiers"}

	dbTracks.Init()
	http.HandleFunc("/snowtrack/track/", trackHandler)
	http.HandleFunc("/snowtrack/skier/", skierHandler)
	http.ListenAndServe(getPort(), nil)

}