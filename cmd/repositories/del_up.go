package repositories

import (
	"cal/cmd/models"
	"cal/db"
	"fmt"
	"log"
)

// function to update any existing holiday in database
func UpdateH(Holiday models.Holiday, id int) (models.Holiday, error) {
	db := db.InitDB()

	sqlStatement := ` UPDATE holidays SET name = $2, iso_date = $3, international = $4  WHERE id = $1 `
	res, err := db.Exec(sqlStatement, id, Holiday.Name, Holiday.Date.ISO, Holiday.International)
	if err != nil {
		return models.Holiday{}, fmt.Errorf("error updating holiday: %w", err)
	}
	// Check if any rows were affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return models.Holiday{}, fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return models.Holiday{}, fmt.Errorf("no holiday found with ID %d", id)
	}

	Holiday.ID = id
	return Holiday, nil
}

// function to delete an existing holiday from database
func DeleteH(isoDate string) error {

	// Get the DB from the db library
	db := db.InitDB()
	if db == nil {
		log.Println("Database connection is nil")
		return fmt.Errorf("database connection is nil")
	}
	defer db.Close()

	// SQL query to delete the holiday
	result, err := db.Exec(`DELETE FROM holidays WHERE iso_date = $1`, isoDate)
	if err != nil {
		return fmt.Errorf("error deleting data: %w", err)
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no holiday found with ISO date %s", isoDate)
	}

	// Successfully deleted
	return nil
}

// function to delete all data or holidays from database
func DelAll() ([]models.Holiday, error) {

	//get the db
	db := db.InitDB()
	if db == nil {
		log.Println("Database connection is nil")
		return nil, fmt.Errorf("database connection is nil")
	}

	// Query to delete all holidays
	rows, err := db.Query(`Delete FROM holidays`)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %w", err)
	}
	defer rows.Close()

	//get successful
	return []models.Holiday{}, nil
}
