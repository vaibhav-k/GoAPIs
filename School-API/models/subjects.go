package models

import (
	"fmt"
	"net/http"

	"../utils"
)

// SubjectAdd adds subjects to their table and also to the classes specified
type SubjectAdd struct {
	SubjectID int    `json:"subject_id"`
	Title     string `json:"title,omitempty"`
	Classes   []int  `json:"classes"`
}

// Unique returns unique items in a slice
func Unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

// GetSubjects gets the subjects for each class from the database
func GetSubjects() ([]string, error) {
	// Query the DB
	s := fmt.Sprintf("SELECT `standard` FROM `school_classes` sc JOIN `school_class_to_subject` scs ON sc.class_id = scs.class_id JOIN `school_subjects` ss ON ss.subject_id = scs.subject_id")
	result, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}

	// Make the names array
	var names []string

	defer result.Close()

	for result.Next() {
		var name string
		result.Scan(&name)
		names = append(names, name)
	}
	uniquenames := Unique(names)

	return uniquenames, nil
}

// AddSubject adds a new notice to the database
func AddSubject(w http.ResponseWriter, r *http.Request, subject SubjectAdd) {
	// Insert into the DB
	s := fmt.Sprintf("INSERT INTO `school_subjects`(`subject_id`, `title`) VALUES (%d, '%s')", subject.SubjectID, subject.Title)
	result, err := db.Query(s)

	if err == nil || result != nil {
		ResponseJSON(w, subject.Title+utils.AddedSomething)
	} else {
		fmt.Println(err)
		ResponseJSON(w, utils.InsertionFailed)
	}

	// Add the subjects to the classes specified
	for class := range subject.Classes {
		t := fmt.Sprintf("INSERT INTO `school_class_to_subject`(`class_id`, `subject_id`) VALUES (%d, %d)", subject.Classes[class], subject.SubjectID)
		insert, er := db.Query(t)

		if er == nil || insert != nil {
			ResponseJSON(w, fmt.Sprintf("%d %s", subject.Classes[class], utils.AddedSomething))
		} else {
			ResponseJSON(w, er)
		}
	}
}
