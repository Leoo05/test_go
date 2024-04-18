package service

type IProfessorService interface {
	GetProfessorByID()
	PostProfessor()
	PutProfessor()
	DeleteProfessor()
}
