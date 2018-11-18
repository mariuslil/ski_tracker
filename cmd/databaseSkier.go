package main

import "fmt"


func (db *SkiersMongoDB) Init() {
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

	err = session.DB(db.DatabaseName).C(db.SkiersCollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func (db *SkiersMongoDB) Get(id int) (Skier, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close().

	skier := Skier{}
	allWasGood := true

	err = session.DB(db.DatabaseName).C(db.SkiersCollectionName).Find(bson.M{"id": id}).One(&skier)
	if err != nil {
		allWasGood = false
	}

	return skier, allWasGood
}


func (db *SkiersMongoDB) Add(s Skier) int {

	session, err := mgo.Dial(db.DatabaseURL)
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


func (db *SkiersMongoDB) Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// handle to "db"
	count, err := session.DB(db.DatabaseName).C(db.SkiersCollectionName).Count()
	if err != nil {
		fmt.Printf("error in Count(): %v", err.Error())
		return -1
	}
	return count
}

func (db *SkiersMongoDB) getField(field string, id int) (string, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	allWasGood := true
	err = session.DB(db.DatabaseName).C(db.SkiersCollectionName).Find(nil).Select(bson.M{field: 1}).One(&field)
	if err != nil {
		allWasGood = false
	}
	return field, allWasGood
}

func (db *SkiersMongoDB) GetAll() []Skier {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var all []Skier{}

	err = session.DB(db.DatabaseName).C(db.).Find(bson.M{}).All(all)
	if err != nil {
		return []Skier{}
	}

	return all
}