package postgreDatabase

import (
	"fmt"

	"github.com/Leoo05/test_go/internal/domain/student/entities"
	"github.com/Leoo05/test_go/internal/infrastructure/secondary/postgreDatabase/models"
)

//This repository class its the one to interact with the database and get the info to response the usecase class

func (adapter *DBAdapter) CreateStudent(student *entities.Student) error {
	modelStudent := &models.Student{
		IDStudent: student.IDStudent,
		Name:      student.Name,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Document:  student.Document,
		Phone:     student.Phone,
		Address:   student.Address,
	}
	err := adapter.conn.Create(&modelStudent).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) ReadStudent(idStudent int) (*entities.Student, error) {
	modelStudent := models.Student{}
	err := adapter.conn.Where("document = ? ", idStudent).Take(&modelStudent).Error
	if err != nil {
		return &entities.Student{}, fmt.Errorf("%s", err.Error())
	}
	entityStudent := &entities.Student{
		IDStudent: modelStudent.IDStudent,
		Name:      modelStudent.Name,
		FirstName: modelStudent.FirstName,
		LastName:  modelStudent.LastName,
		Document:  modelStudent.Document,
		Phone:     modelStudent.Phone,
		Address:   modelStudent.Address,
	}
	return entityStudent, nil
}

func (adapter *DBAdapter) UpdateStudent(idStudent int, student *entities.Student) error {
	modelStudent := &models.Student{
		Name:      student.Name,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Document:  student.Document,
		Phone:     student.Phone,
		Address:   student.Address,
	}
	err := adapter.conn.Model(&models.Student{}).Where("document = ?", idStudent).Updates(&modelStudent).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) DeleteStudent(idStudent int) error {
	err := adapter.conn.Where("document = ?", idStudent).Delete(&models.Student{}).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
