package service

import "github.com/Leoo05/test_go/internal/domain/grade/entities"

type IGradesService interface {
	GetGradeByID(int) (*entities.Grade, error)
	PostGrade(*entities.Grade) error
	PutGrade(int, *entities.Grade) error
	DeleteGrade(int) error
}
