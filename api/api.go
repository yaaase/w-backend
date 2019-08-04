package api

import (
	"airplane-api/storage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Server holds information about a database connection
type Server struct {
	DataRepository *storage.Database
}

// IndexHandler returns a json array of workout days
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	workouts, err := storage.GetWorkouts(s.DataRepository)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(workouts)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

// CreateHandler takes a json POST of workout info and calls to the DB to store it
func (s *Server) CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	var workoutReq storage.WorkoutReq
	err = json.Unmarshal(body, &workoutReq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = s.DataRepository.CreateWorkout(&workoutReq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
