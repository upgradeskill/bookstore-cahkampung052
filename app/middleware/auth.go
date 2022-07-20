package middle

import (
	"bookstore/app/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Auth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func Roles(text string) echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			var roles map[string]map[string]bool

			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// Convert claims from map[string]interface{} to string
			str := claims["roles"].(string)

			// encode json roles
			err := json.Unmarshal([]byte(str), &roles)
			if err != nil {
				fmt.Println(err)
			}

			feature, action := decodeRoleAuth(text)
			if roles[feature][action] {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, pkg.ResponseFormatter("You don't have credential to access this page", nil))
		})
	})
}

func decodeRoleAuth(text string) (feature string, action string) {
	resp := strings.Split(text, ":")
	feature = resp[0]
	action = resp[1]
	return
}
