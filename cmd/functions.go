package main

import "os"

func (db *SnowTrack) Init() {
	db.SnowTracks = make(map[int]SnowTrack)
}

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

func getPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	return ":" + port
}
