package apiRest

import (
	"github.com/labstack/echo"
)

type APIRest struct {
	echoInstance *echo.Echo
	globalGroup  *echo.Group
	port         string
	routes       func(*echo.Group)
}

func NewAPIRest(port, globalPrefix string, routes func(*echo.Group)) *APIRest {
	e := echo.New()
	return &APIRest{
		echoInstance: e,
		globalGroup:  e.Group(globalPrefix),
		port:         port,
		routes:       routes,
	}
}

func (api *APIRest) RunServer() {
	// Indicando la ruta base para la version 1
	api.routes(api.globalGroup)
	api.echoInstance.Logger.Fatal(api.echoInstance.Start(":" + api.port))
}
