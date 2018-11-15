package main
// YOLO
import (
	"net/http"
	"strings"
)

func skierHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body == nil {
			http.Error(w, "POST request must have a JSON body", http.StatusBadRequest)
			return
		}

	case "GET":



	}

}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body == nil {
			http.Error(w, "POST request must have a JSON body", http.StatusBadRequest)
			return
		}

		url := strings.Split(r.URL.Path, "/")


	case "GET":



	}
}