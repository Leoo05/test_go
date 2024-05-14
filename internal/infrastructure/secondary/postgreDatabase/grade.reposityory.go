package postgreDatabase

import (
	"fmt"

	courseEntity "github.com/Leoo05/test_go/internal/domain/course/entities"
	"github.com/Leoo05/test_go/internal/domain/grade/entities"
	studentEntity "github.com/Leoo05/test_go/internal/domain/student/entities"
	"github.com/Leoo05/test_go/internal/infrastructure/secondary/postgreDatabase/models"
)

func (adapter *DBAdapter) CreateGrade(grade *entities.Grade) error {
	modelGrade := &models.Grade{
		IDGrade:   grade.IDGrade,
		IDStudent: uint(grade.IDStudent),
		IDCourse:  uint(grade.IDCourse),
		Grade:     grade.Grade,
	}
	err := adapter.conn.Create(&modelGrade).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) ReadGrade(idGrade int) (*entities.Grade, error) {
	modelGrade := models.Grade{}
	err := adapter.conn.Preload("Course").Preload("Student").Where("id_grade = ?", idGrade).Take(&modelGrade).Error
	if err != nil {
		return &entities.Grade{}, fmt.Errorf("%s", err.Error())
	}
	student := studentEntity.Student{
		Name: modelGrade.Student.Name,
	}
	course := courseEntity.Course{
		CourseName: modelGrade.Course.CourseName,
		CourseCode: modelGrade.Course.CourseCode,
	}
	entityGrade := &entities.Grade{
		IDGrade:   modelGrade.IDGrade,
		IDStudent: int(modelGrade.IDStudent),
		IDCourse:  int(modelGrade.IDCourse),
		Grade:     modelGrade.Grade,
		Student:   student,
		Course:    course,
	}
	return entityGrade, nil
}

func (adapter *DBAdapter) UpdateGrade(idGrade int, grade *entities.Grade) error {
	modelGrade := &entities.Grade{
		IDGrade:   grade.IDGrade,
		IDStudent: grade.IDStudent,
		IDCourse:  grade.IDCourse,
		Grade:     grade.Grade,
	}
	err := adapter.conn.Model(&models.Grade{}).Where("id_grade = ? ", idGrade).Updates(&modelGrade).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) DeleteGrade(idGrade int) error {
	err := adapter.conn.Where("id_grade = ? ", idGrade).Delete(&models.Professor{}).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
