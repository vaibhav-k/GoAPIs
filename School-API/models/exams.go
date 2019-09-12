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

	var exam Exams

	// Make the exam struct
	for result.Next() {
		err := result.Scan(&exam.ExamID, &exam.TypeID, &exam.Title, &exam.Date, &exam.Time)
		if err != nil {
			panic(err.Error())
		}
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

	var exams []Exams

	// Make the exams array
	for result.Next() {
		var exam Exams
		err := result.Scan(&exam.ExamID, &exam.TypeID, &exam.Title, &exam.Date, &exam.Time)
		if err != nil {
			panic(err.Error())
		}
		exams = append(exams, exam)
	}
	return exams, nil
}
