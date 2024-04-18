package repository

type IStudentRepository interface {
	CreateStudent()
	ReadStudent()
	UpdateStudent()
	DeleteStudent()
}
