package snowtrack

import (
	"net/http"
)


var Global_db SnowTrackStorage
var SnowTracks SnowTrack

func main() {

	Global_db = &SnowTracksMongoDB{"mongodb://olebgr:password1@ds261253.mlab.com:61253/snowtracks",
	"snowtracks"}

	Global_db.Init()
	http.HandleFunc("/snowtrack/api", apiHandler)
	http.ListenAndServe(getPort(), nil)

}