package handler

import (
	"net/http"
	"onelab/internal/model"

	"github.com/labstack/echo/v4"
)

func (h Manager) TransactionsCreate(c echo.Context) error {
	var transaction model.TransactionsCreate

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.srv.Transactions.CreateTransaction(c.Request().Context(), transaction); err != nil {
		return c.JSON(http.StatusBadGateway, err)

	}

	return c.JSON(http.StatusOK, "Your transaction is created")
}

func (h Manager) AllTransactions(c echo.Context) error {
	transactions, err := h.srv.Transactions.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transactions)
}
