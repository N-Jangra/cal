package repositories

import (
	"cal/cmd/models"
	"cal/db"
	"fmt"
	"log"
)

// function to create or add an new holiday to database
func AddH(Holiday models.Holiday) (models.Holiday, error) {

	//get the db from db library
	db := db.InitDB()
	if db == nil {
		log.Println("Database connection is nil")
		return Holiday, fmt.Errorf("database connection is nil")
	}

	//extract iso string from db
	isoDate := Holiday.Date.ISO

	//add data in db
	_, err := db.Exec(`insert into holidays (name, iso_date, international) values ($1, $2, $3)`, Holiday.Name, isoDate, Holiday.International)
	if err != nil {
		return Holiday, fmt.Errorf("error inserting data: %w", err)
	}

	//insert successfull
	return Holiday, nil
}

// function to get all available holidays from database
func GetH() ([]models.Holiday, error) {

	//get the db
	db := db.InitDB()
	if db == nil {
		log.Println("Database connection is nil")
		return nil, fmt.Errorf("database connection is nil")
	}

	// Query to get all holidays
	rows, err := db.Query(`SELECT name, iso_date, international FROM holidays`)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %w", err)
	}
	defer rows.Close()

	// Initialize a slice to hold the results
	var holidays []models.Holiday

	// Loop through the rows and scan each row into a Holiday struct
	for rows.Next() {
		var holiday models.Holiday
		// Scan each column into the struct
		if err := rows.Scan(&holiday.Name, &holiday.Date.ISO, &holiday.International); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		// Add the holiday to the result slice
		holidays = append(holidays, holiday)
	}

	// Check for errors after looping through the rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows iteration: %w", err)
	}

	//get successful
	return holidays, nil
}

// function to get an existing holiday from database
func GetTH(isoDate string) ([]models.Holiday, error) {

	// Get the DB from the db library
	db := db.InitDB()
	if db == nil {
		log.Println("Database connection is nil")
		return nil, fmt.Errorf("database connection is nil")
	}
	defer db.Close()

	// Query to get all holidays
	rows, err := db.Query(`SELECT name, iso_date, international FROM holidays WHERE iso_date = $1`, isoDate)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %w", err)
	}
	defer rows.Close()

	// Initialize a slice to hold the results
	var holidays []models.Holiday

	// Loop through the rows and scan each row into a Holiday struct
	for rows.Next() {
		var holiday models.Holiday
		// Scan each column into the struct
		if err := rows.Scan(&holiday.Name, &holiday.Date.ISO, &holiday.International); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		// Add the holiday to the result slice
		holidays = append(holidays, holiday)
	}

	// Check for errors after looping through the rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows iteration: %w", err)
	}

	//get successful
	return holidays, nil
}
