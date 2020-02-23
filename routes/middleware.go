package routes

import (
	"fmt"

	"github.com/labstack/echo"
)

func FirstEncounter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("ini middleware")
		return next(c)
	}
}
