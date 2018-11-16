package main

import (
	"net/http"
)

var db SnowTrackStorage
var SnowTracks SnowTrack

func main() {
	db = &SnowTracksMongoDB{"mongodb://olebgr:password1@ds261253.mlab.com:61253/snowtracks",
	"snowtracks"}

	db.Init()
	http.HandleFunc("/snowtrack/track/", trackHandler)
	http.HandleFunc("/snowtrack/skier/", skierHandler)
	http.ListenAndServe(getPort(), nil)

}