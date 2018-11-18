package main

import "fmt"


func (db *SkiersMongoDB) Init() {
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

	err = session.DB(db.DatabaseName).C(db.SkiersCollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func (db *SkiersMongoDB) Get(id int) (Skier, bool) {
	session, err := mgo.Dial(db.DatabaseName)
	if err != nil {
		panic(err)
	}
	defer session.Close().

		skier := Skier{}
	allWasGood := true
	err = session.DB(db.DatabaseName).C(db.SkiersCollectionName).Find(bson.M{"id": id}).One(&Skier{})
	if err != nil {
		allWasGood = false
	}

	return skier, allWasGood
}


func (db *SkiersMongoDB) Add(s Skier) int {

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

	err = session.DB(db.DatabaseName).C(db.SkiersCollectionName).Insert(s)
	if err!= nil {
		fmt.Printf("error in Insert(): %v", err.Error())
	}
	return id
}


