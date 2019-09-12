package models

import (
	"fmt"
	"net/http"

	"../utils"
)

// Sections struct for sections to IDs
type Sections struct {
	ClassID        int `json:"class_id"`
	SectionID      int `json:"section_id"`
	ClassSectionID int `json:"class_section_id"`
}

// GetSections gets all the details of all sections from the database
func GetSections() ([]Sections, error) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_sections`")
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var sections []Sections

	// Make the sections array
	for result.Next() {
		var section Sections
		result.Scan(&section.ClassID, &section.SectionID, &section.ClassSectionID)
		sections = append(sections, section)
	}

	return sections, nil
}

// AddSection gets all the details of all sections from the database
func AddSection(w http.ResponseWriter, r *http.Request, section Sections) {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_sections` SET `class_id` = %d, `section` = %d, `class_section_id` = %d", section.ClassID, section.SectionID, section.ClassSectionID)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, fmt.Sprintf("%d %s", section.ClassSectionID, utils.AddedSomething))
	} else {
		ResponseJSON(w, err)
	}
}
