package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
}

func (h *UsersHandler) GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "ini endpoint user")
}
