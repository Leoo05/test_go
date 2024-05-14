package repository

import "github.com/Leoo05/test_go/internal/domain/professor/entities"

type IProfessorRepository interface {
	CreateProfessor(*entities.Professor) error
	ReadProfessor(int) (*entities.Professor, error)
	UpdateProfessor(int, *entities.Professor) error
	DeleteProfessor(int) error
}
