package handler

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createMarketPlace(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	var input posCreditation.TodoMarketPlace
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoMarketPlace.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllItemsResponse struct {
	Data []posCreditation.TodoMarketPlace `json:"data"`
}

func (h *Handler) getAllMarketPlaces(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	marketPlaces, err := h.services.TodoMarketPlace.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: marketPlaces,
	})
}

func (h *Handler) getMarketPlaceById(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	marketPlaceId := c.Param("id")

	item, err := h.services.TodoMarketPlace.GetById(marketPlaceId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateMarketPlace(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	marketPlaceId := c.Param("id")

	var input posCreditation.UpdateMarketPlaceInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoMarketPlace.UpdateById(marketPlaceId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) deleteMarketPlace(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
	}

	id := c.Param("id")

	err := h.services.TodoMarketPlace.DeleteById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
