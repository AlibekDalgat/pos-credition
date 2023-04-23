package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	h.roling(c, input)
	var (
		token string
		err   error
		check bool
	)
	check = checkRole(c)
	if !check {
		token, err = h.services.Authorization.GenerateTokenForAgent(input.Login, input.Password)
	} else {
		token, err = h.services.Authorization.GenerateTokenForAdmin()
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		isAdmin: check,
	})
}
