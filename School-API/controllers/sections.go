package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../utils"
)

// GetSections gets all the details of all sections from the database
func GetSections(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "GET" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the GET method for this route")
	}
	fmt.Println("Getting section details!")

	// Call the handler
	sections, err := models.GetSections()

	if err != "" {
		sectiondetails := models.Response{
			StatusCode: http.StatusNoContent,
			Message:    utils.GetFailed,
			Data:       sections,
		}

		// Return from the function
		ResponseJSON(w, sectiondetails)
		return
	}

	sectiondetails := models.Response{
		StatusCode: utils.SuccessCode,
		Message:    utils.GotSections,
		Data:       sections,
	}

	// Return from the function
	ResponseJSON(w, sectiondetails)
}

// AddSection gets all the details of all sections from the database
func AddSection(w http.ResponseWriter, r *http.Request) {
	// Check if the Method is correct
	if r.Method != "POST" {
		http.Error(w, http.StatusText(utils.WrongMethod), utils.WrongMethod)
		ResponseJSON(w, "Please use the POST method for this route")
	}
	fmt.Println("Adding a new section!")

	// Get the user details from the POST body
	decoder := json.NewDecoder(r.Body)
	var section models.Sections
	err := decoder.Decode(&section)
	if err != nil {
		panic(err)
	}

	// Check the user's input and then call the handler
	if section.SectionID < 0 {
		ResponseJSON(w, "Please enter a valid section ID")
	} else if section.ClassID < 0 {
		ResponseJSON(w, "Please enter a valid class ID")
	} else if section.ClassSectionID < 0 {
		ResponseJSON(w, "Please enter a valid class-section ID")
	} else {
		err := models.AddSection(w, r, section)

		if err != "" {
			sectiondetails := models.Response{
				StatusCode: utils.WrongParam,
				Message:    utils.InsertionFailed,
				Data:       "Section is not added!",
			}

			// Return from the function
			ResponseJSON(w, sectiondetails)
			return
		}

		sectiondetails := models.Response{
			StatusCode: utils.SuccessCode,
			Message:    "Section is now added",
			Data:       "Section is now added",
		}

		// Return from the function
		ResponseJSON(w, sectiondetails)
	}
}
