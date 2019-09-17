package models

import (
	"fmt"
	"net/http"
)

// Teachers maps TeacherIDs to information about the teachers
type Teachers struct {
	TeacherID int    `json:"teacher_id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	EmailID   string `json:"email_id,omitempty"`
	Password  string `json:"password,omitempty"`
}

// GetTeacher gets the teacher's detail from the database
func GetTeacher(id string) (Teachers, string) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_teachers` WHERE `teacher_id` = '%s'", id)
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var teacher Teachers
	i := 0

	// Make the teacher struct
	for result.Next() {
		result.Scan(&teacher.TeacherID, &teacher.FirstName, &teacher.LastName, &teacher.EmailID, &teacher.Password)
		i = i + 1
	}

	if i == 0 {
		return teacher, "No teacher with this ID"
	}

	return teacher, ""
}

// GetTeachers gets all teachers detail from the database
func GetTeachers() ([]Teachers, string) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_teachers`")
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var teachers []Teachers

	// Make the teachers array
	for result.Next() {
		var teacher Teachers
		result.Scan(&teacher.TeacherID, &teacher.FirstName, &teacher.LastName, &teacher.EmailID, &teacher.Password)
		teachers = append(teachers, teacher)
	}

	if teachers == nil {
		return teachers, "No teahers right now!"
	}
	return teachers, ""
}

// AddTeacher adds marks to the database
func AddTeacher(w http.ResponseWriter, r *http.Request, teacher Teachers) string {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_teachers` SET `teacher_id` = %d, `first_name` = '%s', `last_name` = '%s', `email_id` = '%s', `password` = '%s'", teacher.TeacherID, teacher.FirstName, teacher.LastName, teacher.EmailID, teacher.Password)
	result, err := db.Query(s)

	if err == nil || result != nil {
		return ""
	}
	return "Could not add the teacher"
}

// DeleteTeacher deletes a teacher from the database
func DeleteTeacher(w http.ResponseWriter, r *http.Request, id string) string {

	s := fmt.Sprintf("SELECT * FROM `school_teachers` WHERE teacher_id = '%s'", id)
	result, _ := db.Query(s)

	// Make the student struct
	var teacher Teachers
	for result.Next() {
		result.Scan(&teacher.TeacherID, &teacher.FirstName, &teacher.LastName, &teacher.EmailID, &teacher.Password)
	}

	if teacher.TeacherID != 0 {
		// Query the DB
		s := fmt.Sprintf("DELETE FROM `school_teachers` WHERE teacher_id = '%s'", id)
		result, err := db.Query(s)

		if err == nil || result != nil {
			return ""
		}
	}
	return "Teacher could not be deleted"
}

// UpdateTeacher updates details of a teacher
func UpdateTeacher(w http.ResponseWriter, r *http.Request, id string, teacher Teachers) string {
	// Check if the ID exists
	s := fmt.Sprintf("SELECT * FROM `school_teachers` WHERE `teacher_id` = '%s'", id)
	result, _ := db.Query(s)

	var teacherScan Teachers
	for result.Next() {
		result.Scan(&teacherScan.TeacherID, &teacherScan.FirstName, &teacherScan.LastName, &teacherScan.EmailID, &teacherScan.Password)
	}

	if teacherScan.TeacherID != 0 {
		// Query the DB
		s := fmt.Sprintf("UPDATE `school_teachers` SET `first_name` = '%s', `last_name` = '%s', `email_id` = '%s', `password` = '%s' WHERE `teacher_id` = '%s'", teacher.FirstName, teacher.LastName, teacher.EmailID, teacher.Password, id)
		_, err := db.Query(s)

		if err == nil {
			return ""
		}
	}
	return "Updating failed!"
}
