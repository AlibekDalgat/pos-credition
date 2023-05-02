package handler

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createAgent(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
		return
	}

	var input posCreditation.TodoAgent
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoAgent.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllAgentsResponse struct {
	Data []posCreditation.TodoAgent `json:"data"`
}
type getInfoMPAgentResponse struct {
	Data []posCreditation.InfoMPsAgent `json:"data"`
}

func (h *Handler) getAllAgents(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
		return
	}

	agents, err := h.services.TodoAgent.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAgentsResponse{
		Data: agents,
	})
}

func (h *Handler) getAgentById(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
		return
	}

	agent := c.Param("id")

	infoMPsAgent, err := h.services.TodoAgent.GetById(agent)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getInfoMPAgentResponse{
		Data: infoMPsAgent,
	})
}

func (h *Handler) updateAgent(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
		return
	}

	agentId := c.Param("id")

	var input posCreditation.UpdateAgentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoAgent.UpdateById(agentId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) deleteAgent(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
		return
	}

	id := c.Param("id")

	err := h.services.TodoAgent.DeleteById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) newAccess(c *gin.Context) {
	if ok := checkRole(c); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "нет прав")
		return
	}

	agentId := c.Param("id")
	var input posCreditation.AccessingToMP
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoAgent.NewAccessToMP(input, agentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
