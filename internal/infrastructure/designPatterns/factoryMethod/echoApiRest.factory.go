package factorymethod

import (
	"os"

	"github.com/Leoo05/test_go/internal/infrastructure/interfaces"
	echoapirest "github.com/Leoo05/test_go/internal/infrastructure/primary/apiRest"
	"github.com/Leoo05/test_go/internal/infrastructure/primary/apiRest/routes"
)

func NewEchoAPIRestAdapter() interfaces.IAPIRest {
	return echoapirest.NewAPIRest(
		os.Getenv("API_REST_PORT"),
		os.Getenv("API_REST_BASE_URL"),
		routes.TestGoRoutes,
	)
}
