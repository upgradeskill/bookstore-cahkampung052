package middle

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Auth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func Roles(text string) echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			header := c.Request().Header
			fmt.Println(header["Authorization"])

			return next(c)
		})
	})
}

// func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
// 		return next(c)
// 	}
// }

// func ServerHeader(text string) echo.MiddlewareFunc {
// 	return echo.MiddlewareFunc(func(h echo.HandlerFunc) echo.HandlerFunc {
// 		return echo.HandlerFunc(func(c echo.Context) error {
// 			return h(c)
// 		})
// 	})
// }
