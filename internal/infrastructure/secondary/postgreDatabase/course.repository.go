package postgreDatabase

import (
	"fmt"

	"github.com/Leoo05/test_go/internal/domain/course/entities"
	"github.com/Leoo05/test_go/internal/infrastructure/secondary/postgreDatabase/models"
)

func (adapter *DBAdapter) CreateCourse(course *entities.Course) error {
	modelCourse := &models.Course{
		IDCourse:   course.IDCourse,
		CourseName: course.CourseName,
		CourseCode: course.CourseCode,
	}
	err := adapter.conn.Create(&modelCourse).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) ReadCourse(idCourse string) (*entities.Course, error) {
	modelCourse := models.Course{}
	err := adapter.conn.Where("course_code = ? ", idCourse).Take(&modelCourse).Error
	if err != nil {
		return &entities.Course{}, fmt.Errorf("%s", err.Error())
	}
	entityCourse := &entities.Course{
		IDCourse:   modelCourse.IDCourse,
		CourseName: modelCourse.CourseName,
		CourseCode: modelCourse.CourseCode,
	}
	return entityCourse, nil
}

func (adapter *DBAdapter) UpdateCourse(idCourse string, course *entities.Course) error {
	modelCourse := &models.Course{
		CourseName: course.CourseName,
		CourseCode: course.CourseCode,
	}
	err := adapter.conn.Model(&models.Course{}).Where("course_code = ? ", idCourse).Updates(&modelCourse).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) DeleteCourse(idCourse string) error {
	err := adapter.conn.Where("course_code = ?", idCourse).Delete(&models.Course{}).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
