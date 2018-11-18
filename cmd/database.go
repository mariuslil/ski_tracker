package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

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

	err = session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func (db *SnowTracksMongoDB) Get(id int) (SnowTrack, bool) {
	session, err := mgo.Dial(db.DatabaseName)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	snowtrack := SnowTrack{}
	allWasGood := true
	err = session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).Find(bson.M{"id": id}).One(&SnowTrack)
	if err != nil {
		allWasGood = false
	}

	return snowtrack, allWasGood
}


func (db *SnowTracksMongoDB) Add(t SnowTrack) int {

	session, err := mgo.Dial(db.DatabaseName)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Creating a id for the track
	id := rand.Int()

	// Checking if track exists in db
	_, ok := db.Get(id)

	// Creating a new id if it exists
	for ok {
		id = rand.Int()
		_, ok = db.Get(id)
	}

	err = session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).Insert(t)
	if err!= nil {
		fmt.Printf("error in Insert(): %v", err.Error())
	}
	return id
}

