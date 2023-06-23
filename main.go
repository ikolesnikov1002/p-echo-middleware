package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/status", Handler, UserRoleMiddleware)

	e.Logger.Fatal(e.Start(":1323"))
}

func Handler(ctx echo.Context) error {
	futureDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()
	dateDiff := futureDate.Sub(currentDate)

	s := fmt.Sprintf("Days left %d", int64(dateDiff.Hours()/24))

	err := ctx.String(http.StatusOK, s)

	if err != nil {
		return err
	}

	return nil
}

func UserRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole := c.Request().Header.Get("User-Role")

		if userRole == "admin" {
			fmt.Println("red button user detected")
		}

		return next(c)
	}
}
