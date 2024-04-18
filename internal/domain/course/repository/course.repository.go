package repository

type ICourseRepository interface {
	CreateCourse()
	ReadCourse()
	UpdateCourse()
	DeleteCourse()
}
