package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"strconv"
)


func (db *SnowTracksMongoDB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
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

func (db *SnowTracksMongoDB) Get(id string) (SnowTrack, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	snowtrack := SnowTrack{}
	allWasGood := true
	err = session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).Find(bson.M{"id": id}).One(snowtrack)
	if err != nil {
		allWasGood = false
	}

	return snowtrack, allWasGood
}


func (db *SnowTracksMongoDB) Add(t SnowTrack) int {

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Creating a id for the track
	id := rand.Int()

	// Checking if track exists in db
	_, ok := db.Get(strconv.Itoa(id))

	// Creating a new id if it exists
	for ok {
		id = rand.Int()
		_, ok = db.Get(strconv.Itoa(id))
	}

	err = session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).Insert(t)
	if err!= nil {
		fmt.Printf("error in Insert(): %v", err.Error())
	}
	return id
}

func (db *SnowTracksMongoDB) Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// handle to "db"
	count, err := session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).Count()
	if err != nil {
		fmt.Printf("error in Count(): %v", err.Error())
		return -1
	}
	return count
}

func (db *SnowTracksMongoDB) getField(field string, id int) (string, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	allWasGood := true
	err = session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).Find(nil).Select(bson.M{field: 1}).One(&field)
	if err != nil {
		allWasGood = false
	}
	return field, allWasGood
}

func (db *SnowTracksMongoDB) GetAll() []SnowTrack {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var all []SnowTrack

	err = session.DB(db.DatabaseName).C(db.SnowTracksCollectionName).Find(bson.M{}).All(all)
	if err != nil {
		return []SnowTrack{}
	}

	return all
}