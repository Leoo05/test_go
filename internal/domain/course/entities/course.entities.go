package entities

type Course struct {
	IDCourse   uint   `json:"idCourse,omitempty"`
	CourseName string `json:"courseName"`
	CourseCode string `json:"courseCode"`
}
