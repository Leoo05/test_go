package routes

import (
	courseService "github.com/Leoo05/test_go/internal/application/usecases/course"
	gradeService "github.com/Leoo05/test_go/internal/application/usecases/grade"
	professorService "github.com/Leoo05/test_go/internal/application/usecases/professor"
	studentService "github.com/Leoo05/test_go/internal/application/usecases/student"
	courseHandler "github.com/Leoo05/test_go/internal/infrastructure/primary/handlers/course"
	gradeHandler "github.com/Leoo05/test_go/internal/infrastructure/primary/handlers/grade"
	professorHandler "github.com/Leoo05/test_go/internal/infrastructure/primary/handlers/professor"
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
	professorServ := professorService.NewProfessorCase(dbAdapter)
	professorApi := professorHandler.NewAPI(professorServ)
	gradeServ := gradeService.NewGradeCase(dbAdapter)
	gradeApi := gradeHandler.NewAPI(gradeServ)
	courseServ := courseService.NewCourseCase(dbAdapter)
	courseApi := courseHandler.NewAPI(courseServ)

	// student api routes
	group.GET("v1/student/:id_student", studentApi.GetStudentByID())
	group.POST("v1/student", studentApi.PostStudent())
	group.PUT("v1/student/:id_student", studentApi.PutStudent())
	group.DELETE("v1/student/:id_student", studentApi.DeleteStudent())

	// professor api routes
	group.GET("v1/professor/:id_professor", professorApi.GetProfessorByID())
	group.POST("v1/professor", professorApi.PostProfessor())
	group.PUT("v1/professor/:id_professor", professorApi.PutProfessor())
	group.DELETE("v1/professor/:id_professor", professorApi.DeleteProfessor())

	// course api routes
	group.GET("v1/course/:id_course", courseApi.GetCourseByID())
	group.POST("v1/course", courseApi.PostCourse())
	group.PUT("v1/course/:id_course", courseApi.PutCourse())
	group.DELETE("v1/course/:id_course", courseApi.DeleteCourse())

	// grade api routes
	group.GET("v1/grade/:id_grade", gradeApi.GetGradeByID())
	group.POST("v1/grade", gradeApi.PostGrade())
	group.PUT("v1/grade/:id_grade", gradeApi.PutGrade())
	group.DELETE("v1/grade/:id_grade", gradeApi.DeleteGrade())
}
