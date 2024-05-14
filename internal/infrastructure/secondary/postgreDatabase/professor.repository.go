package postgreDatabase

import (
	"fmt"

	"github.com/Leoo05/test_go/internal/domain/professor/entities"
	"github.com/Leoo05/test_go/internal/infrastructure/secondary/postgreDatabase/models"
)

func (adapter *DBAdapter) CreateProfessor(professor *entities.Professor) error {
	modelProfessor := &models.Professor{
		IDProfessor: professor.IDProfessor,
		Name:        professor.Name,
		FirstName:   professor.FirstName,
		LastName:    professor.LastName,
		Document:    professor.Document,
		Phone:       professor.Phone,
	}
	err := adapter.conn.Create(&modelProfessor).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) ReadProfessor(idProfessor int) (*entities.Professor, error) {
	modelProfessor := models.Professor{}
	err := adapter.conn.Where("document = ? ", idProfessor).Take(&modelProfessor).Error
	if err != nil {
		return &entities.Professor{}, fmt.Errorf("%s", err.Error())
	}
	entityProfessor := &entities.Professor{
		IDProfessor: modelProfessor.IDProfessor,
		Name:        modelProfessor.Name,
		FirstName:   modelProfessor.FirstName,
		LastName:    modelProfessor.LastName,
		Document:    modelProfessor.Document,
		Phone:       modelProfessor.Phone,
	}
	return entityProfessor, nil
}

func (adapter *DBAdapter) UpdateProfessor(idProfessor int, professor *entities.Professor) error {
	modelProfessor := &models.Professor{
		Name:      professor.Name,
		FirstName: professor.FirstName,
		LastName:  professor.LastName,
		Document:  professor.Document,
		Phone:     professor.Phone,
	}
	err := adapter.conn.Model(&models.Professor{}).Where("document = ?", idProfessor).Updates(&modelProfessor).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (adapter *DBAdapter) DeleteProfessor(idProfessor int) error {
	err := adapter.conn.Where("document = ?", idProfessor).Delete(&models.Professor{}).Error
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
