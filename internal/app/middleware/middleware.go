package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const roleAdmin = "admin"

func UserRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRole := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(userRole, roleAdmin) {
			log.Println("red button user detected")
		}

		return next(ctx)
	}
}
