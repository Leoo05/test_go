package repository

type IGradesRepository interface {
	CreateGrade()
	ReadGrade()
	UpdateGrade()
	DeleteGrade()
}
