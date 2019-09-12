package models

import (
	"fmt"
)

// Exams tell us which exam is when
type Exams struct {
	ExamID int    `json:"exam_id"`
	TypeID int    `json:"type_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Time   string `json:"time"`
}

// ExamTypes tell us which ExamID corresponds to which type of exam
type ExamTypes struct {
	ExamID int    `json:"exam_id"`
	Title  string `json:"title"`
	Year   string `json:"year"`
}

// GetExam gets the teacher's detail from the database
func GetExam(id string) (Exams, error) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_exams` WHERE `exam_id` = '%s'", id)
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the exam struct
	var exam Exams
	for result.Next() {
		result.Scan(&exam.ExamID, &exam.TypeID, &exam.Title, &exam.Date, &exam.Time)
	}
	return exam, nil
}

// GetExams gets all the exams to be held from the database
func GetExams() ([]Exams, error) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_exams`")
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the exams array
	var exams []Exams
	for result.Next() {
		var exam Exams
		result.Scan(&exam.ExamID, &exam.TypeID, &exam.Title, &exam.Date, &exam.Time)
		exams = append(exams, exam)
	}
	return exams, nil
}
