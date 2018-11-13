package snowtrack

import (
	"gopkg.in/mgo.v2"
	"os"
)

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

}

/*func (db *SnowTracks) Init() {
	db.SnowTracks = make(map[int]SnowTrack)
}*/

func (db *SnowTracksMongoDB) Init() {
	session, err := mgo.Dial(db.DatabaseName)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	index := mgo.Index{
		Key: []string{},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}

	err = session.DB(db.DatabaseName).C(db.TracksCollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

var Global_db SnowTrackStorage
var SnowTracks SnowTrack

func main() {

	Global_db = &SnowTracksMongoDB{"mongodb://olebgr:password1@ds261253.mlab.com:61253/snowtracks",
	"snowtracks"}

	Global_db.Init()
	var p string
	if port := os.Getenv("PORT"); port != "" {
		p = ":" + port
	} else {
		p = ":8080"
	}
}
