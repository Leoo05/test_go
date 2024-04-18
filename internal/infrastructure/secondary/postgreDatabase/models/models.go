package models

type Student struct {
	IDStudent uint `gorm:"column:id_student;primaryKey"`
	Name      string
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Document  int
	Phone     int
	Address   string
}

type Course struct {
	IDCourse   uint   `gorm:"column:id_course;primaryKey"`
	CourseName string `gorm:"column:course_name"`
}

type Grade struct {
	IDGrade   uint    `gorm:"column:id_grade;primaryKey"`
	IDStudent uint    `gorm:"column:fk_student"`
	Student   Student `gorm:"foreignKey:IDStudent;references:IDStudent"`
	IDCourse  uint    `gorm:"column:fk_course"`
	Course    Course  `gorm:"foreignKey:IDCourse;references:IDCourse"`
}

type Professor struct {
	IDProfessor uint `gorm:"column:id_professor;primaryKey"`
	Name        string
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	Document    int
	Phone       int
}

func (Student) TableName() string {
	return "student"
}

func (Professor) TableName() string {
	return "professor"
}

func (Course) TableName() string {
	return "course"
}
func (Grade) TableName() string {
	return "grade"
}
