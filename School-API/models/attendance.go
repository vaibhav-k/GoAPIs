// swagger:operation GET /attendances
//
// Returns attendances of all days
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - date: string
// responses:
//   '200':
//     description: The attendances
//     type: json

package models

import (
	"fmt"
)

// Attendance maps student_IDs to their attendance on each day
type Attendance struct {
	StudentID  int `json:"student_id"`
	Attendance int `json:"attendance"`
}

// GetAttendance gets the attendance of a date from the database
func GetAttendance(key string) ([]Attendance, string) {
	// Query the DB to find the date_id
	s := fmt.Sprintf("SELECT `date_id` FROM `school_date_to_id` WHERE date = '%s'", key)
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the attendance array
	var attds []Attendance
	for result.Next() {
		date := 0
		result.Scan(&date)

		// Query the DB to find the attendance on the day
		u := fmt.Sprintf("SELECT `student_id`, `attendance` FROM `school_attendance` e JOIN `school_date_to_id` r ON e.date_id=r.date_id WHERE e.date_id = %d", date)
		result2, err2 := db.Query(u)
		if err2 != nil {
			panic(err.Error())
		}

		defer result2.Close()

		for result2.Next() {
			var attd Attendance
			result2.Scan(&attd.StudentID, &attd.Attendance)
			attds = append(attds, attd)
		}
	}
	if attds != nil {
		return attds, ""
	}
	return attds, "Wrong date"
}

// GetAttendances gets the attendance of all students from the database
func GetAttendances() ([]Attendance, string) {
	// Query the DB to find all attendances
	u := fmt.Sprintf("SELECT `student_id`, `attendance` FROM `school_attendance` e JOIN `school_date_to_id` r ON e.date_id=r.date_id")
	result, err := db.Query(u)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the attendance array
	var attds []Attendance
	for result.Next() {
		var attd Attendance
		result.Scan(&attd.StudentID, &attd.Attendance)
		attds = append(attds, attd)
	}
	if attds != nil {
		return attds, ""
	}
	return attds, "Wrong date"
}
