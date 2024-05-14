package professor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Leoo05/test_go/internal/domain/professor/entities"
	"github.com/Leoo05/test_go/internal/domain/professor/service"
	"github.com/labstack/echo"
)

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type APIRestProfessor struct {
	service service.IProfessorService
}

// The function of this Handler "class" is to interact at api response level get the body/query params to later pass this info to the usecase level, and to deliver the response the user needs
func NewAPI(service service.IProfessorService) *APIRestProfessor {
	return &APIRestProfessor{service: service}
}

func (a *APIRestProfessor) GetProfessorByID() func(c echo.Context) error {
	return func(c echo.Context) error {
		professorID, err := strconv.ParseInt(c.Param("id_professor"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idProfessor invalid",
				Field:   err.Error(),
			})
		}
		professor, err := a.service.GetProfessorByID(int(professorID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		fmt.Println(professor)
		return c.JSON(http.StatusOK, professor)
	}
}

func (a *APIRestProfessor) PostProfessor() func(c echo.Context) error {
	return func(c echo.Context) error {
		professor := &entities.Professor{}
		err := json.NewDecoder(c.Request().Body).Decode(&professor)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PostProfessor(professor)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, "Professor created succesfully")
	}
}

func (a *APIRestProfessor) PutProfessor() func(c echo.Context) error {
	return func(c echo.Context) error {
		professor := &entities.Professor{}
		professorID, err := strconv.ParseInt(c.Param("id_professor"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idProfessor invalid",
				Field:   err.Error(),
			})
		}
		err = json.NewDecoder(c.Request().Body).Decode(&professor)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PutProfessor(int(professorID), professor)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "Professor updated succesfully")
	}
}

func (a *APIRestProfessor) DeleteProfessor() func(c echo.Context) error {
	return func(c echo.Context) error {
		professorID, err := strconv.ParseInt(c.Param("id_professor"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idProfessor invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.DeleteProfessor(int(professorID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "professor deleted sucessfully")
	}
}
