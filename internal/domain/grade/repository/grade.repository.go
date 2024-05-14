package repository

import "github.com/Leoo05/test_go/internal/domain/grade/entities"

type IGradesRepository interface {
	CreateGrade(*entities.Grade) error
	ReadGrade(int) (*entities.Grade, error)
	UpdateGrade(int, *entities.Grade) error
	DeleteGrade(int) error
}
