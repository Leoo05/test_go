package course

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Leoo05/test_go/internal/domain/course/entities"
	"github.com/Leoo05/test_go/internal/domain/course/service"
	"github.com/labstack/echo"
)

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type APIRestCourse struct {
	service service.ICourseService
}

// The function of this Handler "class" is to interact at api response level get the body/query params to later pass this info to the usecase level, and to deliver the response the user needs
func NewAPI(service service.ICourseService) *APIRestCourse {
	return &APIRestCourse{service: service}
}

func (a *APIRestCourse) GetCourseByID() func(c echo.Context) error {
	return func(c echo.Context) error {
		courseID := c.Param("id_course")
		if courseID == "" {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idCourse invalid",
			})
		}
		course, err := a.service.GetCourseByID(courseID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		fmt.Println(course)
		return c.JSON(http.StatusOK, course)
	}
}

func (a *APIRestCourse) PostCourse() func(c echo.Context) error {
	return func(c echo.Context) error {
		course := &entities.Course{}
		err := json.NewDecoder(c.Request().Body).Decode(&course)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PostCourse(course)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, "Course created succesfully")
	}
}

func (a *APIRestCourse) PutCourse() func(c echo.Context) error {
	return func(c echo.Context) error {
		course := &entities.Course{}
		courseID := c.Param("id_course")
		if courseID == "" {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idCourse invalid",
			})
		}
		err := json.NewDecoder(c.Request().Body).Decode(&course)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request body invalid",
				Field:   err.Error(),
			})
		}
		err = a.service.PutCourse(courseID, course)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "Course updated succesfully")
	}
}

func (a *APIRestCourse) DeleteCourse() func(c echo.Context) error {
	return func(c echo.Context) error {
		courseID := c.Param("id_course")
		if courseID == "" {
			return c.JSON(http.StatusBadRequest, &Error{
				Message: "Invalid Request idCourse invalid",
			})
		}
		err := a.service.DeleteCourse(courseID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, "course deleted sucessfully")
	}
}
