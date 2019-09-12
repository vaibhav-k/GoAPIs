package models

import (
	"fmt"
	"net/http"

	"../utils"
)

// Notice struct for notices
type Notice struct {
	TeacherID int    `json:"teacher_id"`
	NoticeID  int    `json:"notice_id"`
	Notice    string `json:"notice"`
}

// GetNotice gets the notices for a student from the database
func GetNotice(w http.ResponseWriter, r *http.Request, id string) (Notice, error) {
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
		err := result.Scan(&notice.TeacherID, &notice.NoticeID, &notice.Notice)
		if err != nil {
			panic(err.Error())
		}
	}
	return notice, nil
}

// GetNotices gets the notices for a student from the database
func GetNotices() ([]Notice, error) {
	// Query the DB
	s := fmt.Sprintf("SELECT * FROM `school_notices`")
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var notices []Notice

	// Make the notices struct
	for result.Next() {
		var notice Notice
		err := result.Scan(&notice.TeacherID, &notice.NoticeID, &notice.Notice)
		if err != nil {
			panic(err.Error())
		}
		notices = append(notices, notice)
	}
	return notices, nil
}

// AddNotice adds a new notice to the database
func AddNotice(notice Notice) error {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_notices`(`teacher_id`, `notice`) VALUES (%d, '%s')", notice.TeacherID, notice.Notice)
	result, err := db.Query(s)

	if err == nil || result != nil {
		fmt.Println(notice.Notice, utils.AddedSomething)
	} else {
		fmt.Println(utils.InsertionFailed)
		fmt.Println(err)
	}
	return err
}
