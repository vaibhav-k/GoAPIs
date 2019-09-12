package utils

// File with all the constants used in the API

// ErrorTeacher for when only a teacher can access the function
const ErrorTeacher string = "Sorry, this is only available for teachers!"

// ErrorStudent for when only a student can access the function
const ErrorStudent string = "Sorry, this is only available for students!"

// ErrorAdmin for when only an admin can access the function
const ErrorAdmin string = "Sorry, this is only available for admins!"

// NeedToLogIn for when the user first needs to log in
const NeedToLogIn string = "Please log in first!"

// LoggedIn for when the user logs in successfully
const LoggedIn string = "Congratulations, you logged in!"

// WrongCreds for when the user enters the wrong credentials
const WrongCreds string = "Your credentials are wrong, please try again!"

// GotNotice for when the notice is presented successfully
const GotNotice string = "Notice below!"

// GotMarks for when the marks are presented successfully
const GotMarks string = "Marks below!"

// GotDetails for when the details are presented successfully
const GotDetails string = "details below"

// GotAttendances for when the attendances are presented successfully
const GotAttendances string = "Attendances below"

// GotSections for when the sections are presented successfully
const GotSections string = "Sections below"

// GotExams for when the exam details are presented successfully
const GotExams string = "Exams below"

// GotTeachers for when the teacher details are presented successfully
const GotTeachers string = "Teachers below"

// GotTeacher for when the teacher details are presented successfully
const GotTeacher string = "Teacher below"

// GotSubjectIDs for when the subject IDs are retrieved successfully
const GotSubjectIDs string = "Subject IDs below"

// GotStudent for when the student details are retrieved successfully
const GotStudent string = "Student details below"

// GetFailed for when the notice is presented successfully
const GetFailed string = "Getting failed!"

// AddedSomething for when the objects are added successfully
const AddedSomething string = "is now added!"

// InsertionFailed for when the insertion fails
const InsertionFailed string = "Insertion failed!"

// DeletedSomething for when the objects are deleted successfully
const DeletedSomething string = "is now deleted!"

// DeletionFailed for when the deletion fails
const DeletionFailed string = "Deletion failed!"

// UpdatedSomething for when the objects are updated successfully
const UpdatedSomething string = "is now updated!"

// UpdatingFailed for when the updating fails
const UpdatingFailed string = "Updating failed!"

// NoType when the user does not specify a type
const NoType string = "Please enter a valid type"

// SuccessCode for when the operation succeeds
const SuccessCode int = 200

// ErrorCode for when the operation fails
const ErrorCode int = 500

// WrongParam for when the operation succeeds
const WrongParam int = 204

// WrongMethod for when the method used is wrong
const WrongMethod int = 500
