package handler

import (
	posCreditation "github.com/AlibekDalgat/pos-credition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCredit(c *gin.Context) {
	agentId, err := getUserId(c)
	mpId := c.Param("mpId")
	var input posCreditation.NewCredit
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Credit.Create(input, mpId, agentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
