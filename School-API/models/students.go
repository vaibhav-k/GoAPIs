package models

import (
	"fmt"
	"net/http"
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
func GetStudents() ([]Students, string) {
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

	if students == nil {
		return students, "No students right now"
	}
	return students, ""
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
func DeleteStudent(w http.ResponseWriter, r *http.Request, id string) string {
	s := fmt.Sprintf("SELECT * FROM `school_students` WHERE student_id = '%s'", id)
	result, _ := db.Query(s)

	// Make the student struct
	var student Students
	for result.Next() {
		result.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.EmailID, &student.Password, &student.Telephone, &student.ClassSectionID)
	}

	if student.StudentID != 0 {
		// Query the DB
		s := fmt.Sprintf("DELETE FROM `school_students` WHERE student_id = '%s'", id)
		result, err := db.Query(s)

		if err == nil || result != nil {
			return ""
		}
	}
	return "Deletion failed!"
}

// UpdateStudent updates details of a student
func UpdateStudent(w http.ResponseWriter, r *http.Request, id string, student Students) string {
	// Check if the ID exists
	s := fmt.Sprintf("SELECT * FROM `school_students` WHERE `student_id` = '%s'", id)
	result, _ := db.Query(s)

	var studentScan Students
	for result.Next() {
		result.Scan(&studentScan.StudentID, &studentScan.FirstName, &studentScan.LastName, &studentScan.EmailID, &studentScan.Password, &studentScan.Telephone, &studentScan.ClassSectionID)
	}

	if studentScan.StudentID != 0 {
		// Query the DB
		s := fmt.Sprintf("UPDATE `school_students` SET `first_name` = '%s', `last_name` = '%s', `email_id` = '%s', `password` = '%s', `telephone` = '%s', `class_section_id` = %d WHERE student_id = '%s'", student.FirstName, student.LastName, student.EmailID, student.Password, student.Telephone, student.ClassSectionID, id)
		result, err := db.Query(s)

		if err == nil && result != nil {
			return ""
		}
	}
	return "Updating failed!"
}
