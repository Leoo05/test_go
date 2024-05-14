package service

import "github.com/Leoo05/test_go/internal/domain/professor/entities"

type IProfessorService interface {
	GetProfessorByID(int) (*entities.Professor, error)
	PostProfessor(*entities.Professor) error
	PutProfessor(int, *entities.Professor) error
	DeleteProfessor(int) error
}
