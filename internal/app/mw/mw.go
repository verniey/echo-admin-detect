package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		val := c.Request().Header.Get("User-Role")

		if strings.EqualFold(val, roleAdmin) {
			log.Println("Red button user detected")
		}

		err := next(c)

		if err != nil {
			return err
		}

		return nil

	}
}
