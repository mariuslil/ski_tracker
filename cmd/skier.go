package main

import (
	"time"
)

type Skier struct {
	SkierID string `json:"skierID"`
	Name string `json:"name"`
	TotalTime time.Duration `json:"totalTime"`
	TotalLength float32 `json:"totalLength"`
	Tracks []SnowTrack `json:"tracks"`
}

type SkierService interface {
	AddSkier(s *Skier) error
	GetSkier(name string) (*Skier, error)
	AddTrack(skier string, track string) error
	AddRide(skier string, track string, time time.Duration, length int) error
}

type SkiersMongoDB struct {
	DatabaseURL string
	DatabaseName string
	SkiersCollectionName string
}

type SkiersStorage interface {
	Init()
	Get(id string) (Skier, bool)
	Add(s Skier) int
	Count() int
	getField(field string, id int) (string, bool)
	GetAll() []Skier
}

type SkierDB struct {
	skiers map[string]Skier
}

func (db *SkierDB) Init() {
	db.skiers = make(map[string]Skier)
}

func (db *SkierDB) Add(S Skier) error {
	db.skiers[S.SkierID] = S
	return nil
}

func (db *SkierDB) Count() int {
	return len(db.skiers)
}

func (db *SkierDB) Get(ID string) (Skier, bool) {
	s, ok := db.skiers[ID]
	return s, ok
}

func (db *SkierDB) GetAll() []Skier {
	all := make([]Skier, 0, db.Count())
	for _, s := range db.skiers {
		all = append(all, s)
	}
}


