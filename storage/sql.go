package storage

import (
	"database/sql"
	// Needed to connect to postgres
	_ "github.com/lib/pq"
)

const (
	getWorkoutsQuery   = "SELECT location, distance, vertical, date from workouts"
	createWorkoutQuery = "INSERT INTO workouts (location, distance, vertical, date) VALUES ($1, $2, $3, $4)"
)

// Database wraps sql.DB
type Database struct {
	DB *sql.DB
}

// Workout is struct wrapping a workouts db row
type Workout struct {
	Location string
	Distance string
	Vertical string
	Date     string
}

// WorkoutReq represents a json POST body
type WorkoutReq struct {
	Location string `json:"location"`
	Distance string `json:"distance"`
	Vertical string `json:"vertical"`
	Date     string `json:"date"`
}

// WorkoutDays is a collection of workout rows
type WorkoutDays struct {
	Days []Workout
}

// InitDB establishes a connection to the database and must be called in main()
func InitDB() (*Database, error) {
	sql, err := sql.Open("postgres", "user=mark dbname=workouts sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = sql.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{DB: sql}, nil
}

// GetWorkouts gets all workout rows in the database, and returns an error if unsuccessful
func GetWorkouts(db *Database) (*WorkoutDays, error) {
	wds := &WorkoutDays{}

	rows, err := db.DB.Query(getWorkoutsQuery)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		w := Workout{}
		err = rows.Scan(
			&w.Location,
			&w.Distance,
			&w.Vertical,
			&w.Date,
		)
		if err != nil {
			return nil, err
		}
		wds.Days = append(wds.Days, w)
	}

	return wds, nil
}

// CreateWorkout writes a workout to the databse and returns an error if unsuccessful
func (db *Database) CreateWorkout(w *WorkoutReq) error {
	statement, err := db.DB.Prepare(createWorkoutQuery)
	if err != nil {
		return err
	}

	_, err = statement.Exec(w.Location, w.Distance, w.Vertical, w.Date)
	if err != nil {
		return err
	}

	return nil
}
