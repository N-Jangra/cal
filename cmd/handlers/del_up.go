package handlers

import (
	"cal/cmd/models"
	"cal/cmd/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// handler for UpdateH function
func Up(c echo.Context) error {
	id := c.Param("id")

	// Convert the ID from string to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// bind the request body to holiday structs
	user := models.Holiday{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// call updateH
	updatedUser, err := repositories.UpdateH(user, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// handler for DeleteH function
func Del(c echo.Context) error {
	// Extract the iso_date from the URL parameter
	isoDate := c.Param("iso_date")
	if isoDate == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ISO date cannot be empty"})
	}

	// Call the DeleteH function to delete the holiday from the database
	err := repositories.DeleteH(isoDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Holiday with ISO date %s has been deleted", isoDate)})
}

// handler for DellAll function
func DelA(c echo.Context) error {

	//call the GetH function to get all holidays
	holidays, err := repositories.DelAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, holidays)
}
