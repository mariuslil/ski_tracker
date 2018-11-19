package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)



func getSkiers (w *http.ResponseWriter, db SkiersMongoDB) {
	if db.Count() == 0 {
		json.NewEncoder(*w).Encode([]Skier{})
	} else {
		a := make([]Skier, 0, db.Count())
		for _, s := range db.GetAll() {
			a = append(a, s)
		}
		json.NewEncoder(*w).Encode(a)
	}

}

func getSkier (w *http.ResponseWriter, db SkiersMongoDB, id string) {
	skier, ok := db.Get(id)
	if !ok {
		http.Error(*w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(*w).Encode(skier)

}

func getTracks (w *http.ResponseWriter, db SnowTracksMongoDB) {
		if db.Count() == 0 {
			json.NewEncoder(*w).Encode(SnowTrack{})
		} else {
			a := make([]SnowTrack, 0, db.Count())
			for _, s := range db.GetAll() {
				a = append(a, s)
			}
			json.NewEncoder(*w).Encode(a)
		}
}

func getTrack (w *http.ResponseWriter, db SnowTracksMongoDB, id string) {
	track, ok := db.Get(id)
	if !ok {
		http.Error(*w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(*w).Encode(track)
}


func skierHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body == nil {
			http.Error(w, "POST request must have a JSON body", http.StatusBadRequest)
			return
		}
		var S Skier
		err := json.NewDecoder(r.Body).Decode(&S)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, ok := dbSkiers.Get(S.SkierID)
		if ok {
			http.Error(w, "Skier already exists", http.StatusBadRequest)
			return
		}

		dbSkiers.Add(S)
		fmt.Fprint(w, "OK")
		return

	case "GET":
		url := strings.Split(r.URL.Path, "/")
		http.Header.Add(w.Header(), "content-type", "text/plain")

		if url[3] == "skiers" {
			getSkiers(&w, SkiersMongoDB{})
		} else if url[3] == "skier" && len(url) == 4 {
			getSkier(&w, SkiersMongoDB{}, url[4])
		} else {http.Error(w, "Bad request, wrong url", http.StatusBadRequest)}

	}

}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body == nil {
			http.Error(w, "POST request must have a JSON body", http.StatusBadRequest)
			return
		}
		var T SnowTrack
		err := json.NewDecoder(r.Body).Decode(&T)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, ok := dbSkiers.Get(T.TrackID)
		if ok {
			http.Error(w, "Track already exists", http.StatusBadRequest)
			return
		}

		dbTracks.Add(T)
		fmt.Fprint(w, "OK")
		return


	case "GET":
		url := strings.Split(r.URL.Path, "/")
		http.Header.Add(w.Header(), "content-type", "text/plain")

		if url[3] == "tracks" {
			getTracks(&w, SnowTracksMongoDB{})
		} else if url[3] == "track" && len(url) == 4 {
			getTrack(&w, SnowTracksMongoDB{}, url[4])
		}


	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")

	if r.Method == "GET" {
		if url[2] == "api" {
			GetAPIinfo(&w)
		} else {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
	}

}