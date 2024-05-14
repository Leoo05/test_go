package entities

import (
	courseEntity "github.com/Leoo05/test_go/internal/domain/course/entities"
	studentEntity "github.com/Leoo05/test_go/internal/domain/student/entities"
)

type Grade struct {
	IDGrade   uint                  `json:"idGrade"`
	IDStudent int                   `json:"idStudent"`
	IDCourse  int                   `json:"idCourse"`
	Grade     int                   `json:"grade"`
	Student   studentEntity.Student `json:"student,omitempty"`
	Course    courseEntity.Course   `json:"course,omitempty"`
}
