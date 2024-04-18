package repository

type IProfessorRepository interface {
	CreateProfessor()
	ReadProfessor()
	UpdateProfessor()
	DeleteProfessor()
}
