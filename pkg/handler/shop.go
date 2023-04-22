package handler

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createShop(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}
	var input posCreditation.TodoShop
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.services.TodoShop.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllShopsResponse struct {
	Data []posCreditation.TodoShop `json:"data"`
}

func (h *Handler) getAllShops(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	lists, err := h.services.TodoShop.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllShopsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "tüzsüz id param")
		return
	}

	list, err := h.services.TodoShop.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateShop(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	id := c.Param("id")
	var input posCreditation.UpdateShopInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoShop.UpdateById(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) deleteShop(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	id := c.Param("id")

	err := h.services.TodoShop.DeleteById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
