package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"onelab/internal/model"

	"github.com/labstack/echo/v4"
)

func (h Manager) CreateUser(c echo.Context) error {
	var user model.UserCreate

	if err := c.Bind(&user); err != nil {
		return err
	}

	err := h.srv.User.Create(c.Request().Context(), user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, "User successfuly created")
}

func (h Manager) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid 'id' parametr")
	}

	user, err := h.srv.User.Get(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}
