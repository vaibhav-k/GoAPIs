package models

import (
	"fmt"
	"net/http"
)

// Sections struct for sections to IDs
type Sections struct {
	ClassID        int `json:"class_id"`
	SectionID      int `json:"section_id"`
	ClassSectionID int `json:"class_section_id"`
}

// GetSections gets all the details of all sections from the database
func GetSections() ([]Sections, string) {
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

	if sections == nil {
		return sections, "Couldn't get the sections"
	}
	return sections, ""
}

// AddSection gets all the details of all sections from the database
func AddSection(w http.ResponseWriter, r *http.Request, section Sections) string {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_sections` SET `class_id` = %d, `section` = %d, `class_section_id` = %d", section.ClassID, section.SectionID, section.ClassSectionID)
	result, err := db.Query(s)

	if err == nil && result != nil {
		return ""
	}
	return "Could not add the section"
}
