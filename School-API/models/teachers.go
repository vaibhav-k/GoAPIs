package models

import (
	"fmt"
	"net/http"

	"../utils"
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
		err := result.Scan(&teacher.TeacherID, &teacher.FirstName, &teacher.LastName, &teacher.EmailID, &teacher.Password)
		if err != nil {
			panic(err.Error())
		}
		i = i + 1
	}

	if i == 0 {
		return teacher, "No teacher with this ID"
	}

	return teacher, ""
}

// GetTeachers gets all teachers detail from the database
func GetTeachers() ([]Teachers, error) {
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
		err := result.Scan(&teacher.TeacherID, &teacher.FirstName, &teacher.LastName, &teacher.EmailID, &teacher.Password)
		if err != nil {
			panic(err.Error())
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

// AddTeacher adds marks to the database
func AddTeacher(w http.ResponseWriter, r *http.Request, teacher Teachers) {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_teachers` SET `teacher_id` = %d, `first_name` = '%s', `last_name` = '%s', `email_id` = '%s', `password` = '%s'", teacher.TeacherID, teacher.FirstName, teacher.LastName, teacher.EmailID, teacher.Password)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, fmt.Sprintf("%d %s", teacher.TeacherID, utils.AddedSomething))
	} else {
		ResponseJSON(w, err)
	}
}

// DeleteTeacher deletes a teacher from the database
func DeleteTeacher(w http.ResponseWriter, r *http.Request, id string) {
	// Query the DB
	s := fmt.Sprintf("DELETE FROM `school_teachers` WHERE teacher_id = '%s'", id)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, fmt.Sprintf("Teacher ID %s %s", id, utils.DeletedSomething))
	} else {
		ResponseJSON(w, err)
	}
}

// UpdateTeacher updates details of a teacher
func UpdateTeacher(w http.ResponseWriter, r *http.Request, id string, teacher Teachers) {
	// Query the DB
	s := fmt.Sprintf("UPDATE `school_teachers` SET `teacher_id` = %d, `first_name` = '%s', `last_name` = '%s', `email_id` = '%s', `password` = '%s' WHERE teacher_id = '%s'", teacher.TeacherID, teacher.FirstName, teacher.LastName, teacher.EmailID, teacher.Password, id)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, fmt.Sprintf("Teacher %s %s", teacher.FirstName, utils.UpdatedSomething))
	} else {
		ResponseJSON(w, err)
	}
}
