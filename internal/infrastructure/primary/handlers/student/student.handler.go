package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Leoo05/test_go/internal/domain/student/entities"
	"github.com/Leoo05/test_go/internal/domain/student/service"
	"github.com/labstack/echo"
)

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type APIRestStudent struct {
	service service.IStudentService
}

// The function of this Handler "class" is to interact at api response level get the body/query params to later pass this info to the usecase level, and to deliver the response the user needs
func NewAPI(service service.IStudentService) *APIRestStudent {
	return &APIRestStudent{service: service}
}

func (a *APIRestStudent) GetStudentByID() func(c echo.Context) error {
	return func(c echo.Context) error {
		studentID, err := strconv.ParseInt(c.Param("id_student"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idStudent invalid",
				Field:   err.Error(),
			})
		}
		student, err := a.service.GetStudent(int(studentID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		fmt.Println(student)
		return c.JSON(http.StatusOK, student)
	}
}

func (a *APIRestStudent) PostStudent() func(c echo.Context) error {
	return func(c echo.Context) error {
		student := &entities.Student{}
		err := json.NewDecoder(c.Request().Body).Decode(&student)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PostStudent(student)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, "Student created succesfully")
	}
}

func (a *APIRestStudent) PutStudent() func(c echo.Context) error {
	return func(c echo.Context) error {
		student := &entities.Student{}
		studentID, err := strconv.ParseInt(c.Param("id_student"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idStudent invalid",
				Field:   err.Error(),
			})
		}
		err = json.NewDecoder(c.Request().Body).Decode(&student)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PutStudent(int(studentID), student)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "Student updated succesfully")
	}
}

func (a *APIRestStudent) DeleteStudent() func(c echo.Context) error {
	return func(c echo.Context) error {
		studentID, err := strconv.ParseInt(c.Param("id_student"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idStudent invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.DeleteStudent(int(studentID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "student deleted sucessfully")
	}
}
