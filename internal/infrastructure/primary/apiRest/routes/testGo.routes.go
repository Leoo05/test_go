package routes

import (
	studentService "github.com/Leoo05/test_go/internal/application/usecases/student"
	studentHandler "github.com/Leoo05/test_go/internal/infrastructure/primary/handlers/student"
	"github.com/Leoo05/test_go/internal/infrastructure/secondary/postgreDatabase"
	"github.com/labstack/echo"
)

func TestGoRoutes(group *echo.Group) {
	dbInstance, err := postgreDatabase.InitDB()
	if err != nil {
		panic(err)
	}
	dbAdapter := postgreDatabase.New(dbInstance)
	studentServ := studentService.NewStudentCase(dbAdapter)
	studentApi := studentHandler.NewAPI(studentServ)
	group.GET("v1/student/:id_student", studentApi.GetStudentByID())
	group.POST("v1/student", studentApi.PostStudent())
	group.PUT("v1/student/:id_student", studentApi.PutStudent())
	group.DELETE("v1/student/:id_student", studentApi.DeleteStudent())
}
