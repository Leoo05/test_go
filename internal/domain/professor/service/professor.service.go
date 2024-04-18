package service

type IProfessorService interface {
	GetProfessorByID()
	PostProfessor()
	PatchProfessor()
	DeleteProfessor()
}
