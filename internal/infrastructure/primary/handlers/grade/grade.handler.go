package grade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Leoo05/test_go/internal/domain/grade/entities"
	"github.com/Leoo05/test_go/internal/domain/grade/service"
	"github.com/labstack/echo"
)

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type APIRestGrade struct {
	service service.IGradesService
}

func NewAPI(service service.IGradesService) *APIRestGrade {
	return &APIRestGrade{service: service}
}

func (a *APIRestGrade) GetGradeByID() func(c echo.Context) error {
	return func(c echo.Context) error {
		gradeID, err := strconv.ParseInt(c.Param("id_grade"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idGrade invalid",
				Field:   err.Error(),
			})
		}
		grade, err := a.service.GetGradeByID(int(gradeID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		fmt.Println(grade)
		return c.JSON(http.StatusOK, grade)
	}
}

func (a *APIRestGrade) PostGrade() func(c echo.Context) error {
	return func(c echo.Context) error {
		grade := &entities.Grade{}
		err := json.NewDecoder(c.Request().Body).Decode(&grade)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PostGrade(grade)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, "Grade created succesfully")
	}
}

func (a *APIRestGrade) PutGrade() func(c echo.Context) error {
	return func(c echo.Context) error {
		grade := &entities.Grade{}
		gradeID, err := strconv.ParseInt(c.Param("id_grade"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idGrade invalid",
				Field:   err.Error(),
			})
		}
		err = json.NewDecoder(c.Request().Body).Decode(&grade)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PutGrade(int(gradeID), grade)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "Grade updated succesfully")
	}
}

func (a *APIRestGrade) DeleteGrade() func(c echo.Context) error {
	return func(c echo.Context) error {
		gradeID, err := strconv.ParseInt(c.Param("id_grade"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idGrade invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.DeleteGrade(int(gradeID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "grade deleted sucessfully")
	}
}
