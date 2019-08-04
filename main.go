package main

import (
	"airplane-api/api"
	"airplane-api/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	fmt.Println("Connected to database.")

	s := &api.Server{
		DataRepository: db,
	}

	http.HandleFunc("/api/workouts", s.IndexHandler)
	http.HandleFunc("/api/workouts/create", s.CreateHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
