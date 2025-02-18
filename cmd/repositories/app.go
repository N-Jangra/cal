package repositories

import (
	"cal/cmd/models"
	"cal/db"
	"fmt"
)

// InsertHoliday inserts a holiday record into the database
func InsertHoliday(holiday models.Holiday) error {
	// SQL query to insert holiday into the database
	query := `INSERT INTO holidays (name, iso_date, international) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, holiday.Name, holiday.Date.ISO, holiday.International)
	if err != nil {
		return fmt.Errorf("failed to insert holiday: %v", err)
	}

	fmt.Printf("Successfully inserted holiday: %s\n", holiday.Name)
	return nil
}
