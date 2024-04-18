package service

type IStudentService interface {
	GetStudentByID()
	PostStudent()
	PatchStudent()
	DeleteStudent()
}
