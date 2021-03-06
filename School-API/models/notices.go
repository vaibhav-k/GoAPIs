package models

import (
	"fmt"
	"net/http"
)

// Notice struct for notices
type Notice struct {
	TeacherID int    `json:"teacher_id"`
	NoticeID  int    `json:"notice_id"`
	Notice    string `json:"notice"`
}

// GetNotice gets the notices for a student from the database
func GetNotice(w http.ResponseWriter, r *http.Request, id string) (Notice, string) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_notices` WHERE `notice_id` = %s", id)
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the notice struct
	var notice Notice
	for result.Next() {
		result.Scan(&notice.TeacherID, &notice.NoticeID, &notice.Notice)
	}

	if notice.NoticeID == 0 {
		return notice, "Wrong ID"
	}

	return notice, ""
}

// GetNotices gets the notices for a student from the database
func GetNotices() ([]Notice, string) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_notices`")
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	// Make the notices struct
	var notices []Notice
	for result.Next() {
		var notice Notice
		result.Scan(&notice.TeacherID, &notice.NoticeID, &notice.Notice)
		notices = append(notices, notice)
	}

	if notices == nil {
		return notices, "No notice right now"
	}
	return notices, ""
}

// AddNotice adds a new notice to the database
func AddNotice(notice Notice) string {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_notices`(`teacher_id`, `notice`) VALUES (%d, '%s')", notice.TeacherID, notice.Notice)
	result, err := db.Query(s)

	if err == nil || result != nil {
		return ""
	}
	return "Notice could not be added"
}
