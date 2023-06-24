package app

import (
	"echo-middleware/internal/app/endpoint"
	"echo-middleware/internal/app/middleware"
	"echo-middleware/internal/app/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.GET("/status", a.e.Status, middleware.UserRole)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	err := a.echo.Start(":1323")

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
