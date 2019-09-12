package models

import (
	"fmt"
	"net/http"

	"../utils"
)

// Students maps StudentIDs to information about the students
type Students struct {
	StudentID      int    `json:"student_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	EmailID        string `json:"email_id"`
	Password       string `json:"password"`
	Telephone      string `json:"telephone"`
	ClassSectionID int    `json:"class_section_id"`
}

// GetStudents gets all of all students from the database
func GetStudents() ([]Students, error) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_students`")
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the students array
	var students []Students
	for result.Next() {
		var student Students
		result.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.EmailID, &student.Password, &student.Telephone, &student.ClassSectionID)
		students = append(students, student)
	}
	return students, nil
}

// GetStudent gets all the details of a student from the database
func GetStudent(id string) (Students, string) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_students` WHERE `student_id` = %s", id)
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the student struct
	var student Students
	i := 0
	for result.Next() {
		result.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.EmailID, &student.Password, &student.Telephone, &student.ClassSectionID)
		i = i + 1
	}
	if i == 0 {
		return student, "No student with this ID"
	}
	return student, ""
}

// DeleteStudent deletes a student from the database
func DeleteStudent(w http.ResponseWriter, r *http.Request, id string) {
	// Query the DB
	s := fmt.Sprintf("DELETE FROM `school_students` WHERE student_id = '%s'", id)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, fmt.Sprintf("Student id %s %s", id, utils.DeletedSomething))
	} else {
		ResponseJSON(w, err)
	}
}

// UpdateStudent updates details of a student
func UpdateStudent(w http.ResponseWriter, r *http.Request, id string, student Students) {
	// Query the DB
	s := fmt.Sprintf("UPDATE `school_students` SET `student_id` = %d, `first_name` = '%s', `last_name` = '%s', `email_id` = '%s', `password` = '%s', `telephone` = '%s', `class_section_id` = %d WHERE student_id = '%s'", student.StudentID, student.FirstName, student.LastName, student.EmailID, student.Password, student.Telephone, student.ClassSectionID, id)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, fmt.Sprintf("Student %s %s", student.FirstName, utils.UpdatedSomething))
	} else {
		ResponseJSON(w, err)
	}
}
