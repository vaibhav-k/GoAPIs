package models

import (
	"fmt"
	"net/http"

	"../utils"
)

// Marks struct for marks
type Marks struct {
	MarksID    int    `json:"marks_id"`
	ExamTypeID int    `json:"exam_type_id,omitempty"`
	Subject    string `json:"subject,omitempty"`
	StudentID  int    `json:"student_id,omitempty"`
	Marks      int    `json:"marks,omitempty"`
}

// GetMarks gets the marks of a student from the database
func GetMarks(id string) ([]Marks, error) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_marks` WHERE `student_id` = %s", id)
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var marks []Marks

	// Make the marks array
	for result.Next() {
		var mark Marks
		err := result.Scan(&mark.MarksID, &mark.ExamTypeID, &mark.Subject, &mark.StudentID, &mark.Marks)
		if err != nil {
			panic(err.Error())
		}
		marks = append(marks, mark)
	}
	return marks, nil
}

// AddMarks adds marks to the database
func AddMarks(w http.ResponseWriter, r *http.Request, mark Marks) {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_marks` SET `exam_type_id` = %d, `subject` = '%s', `student_id` = %d, `marks` = %d", mark.ExamTypeID, mark.Subject, mark.StudentID, mark.Marks)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, "Marks added!")
	} else {
		ResponseJSON(w, utils.InsertionFailed)
	}
}

// UpdateMarks updates marks of an exam
func UpdateMarks(w http.ResponseWriter, r *http.Request, id string, mark Marks) {
	// Query the DB
	s := fmt.Sprintf("UPDATE `school_marks` SET `marks_id` = %d, `exam_type_id` = %d, `subject` = '%s', `student_id` = %d, `marks` = %d WHERE `marks_id` = '%s'", mark.MarksID, mark.ExamTypeID, mark.Subject, mark.StudentID, mark.Marks, id)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, "Marks updated!")
	} else {
		ResponseJSON(w, utils.UpdatingFailed)
	}
}
