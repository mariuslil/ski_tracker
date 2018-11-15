package main

import (
	"net/http"
	"time"
)

type Skier struct {
	Name string `json:"name"`
	TotalTime time.Duration `json:"totalTime"`
	TotalLength float32 `json:"totalLength"`
	Tracks []SnowTrack `json:"tracks"`
}

type SkiersDB struct {
	skiers map[string]Skier
}


func (db *SkiersDB) Init() {
	db.skiers = make(map[string]Skier)
}

func (db *SkiersDB) Add(w http.ResponseWriter, s Skier) string {
	if !exists(s.Name) {
		db.skiers[s.Name] = s
	} else {
		http.Error(w, "Skier allready exists", http.StatusBadRequest)
	}
}

func (db *SkiersDB) Get(n string) (Skier, bool) {
	s, ok := db.skiers[n]
	return s, ok
}