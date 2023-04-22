package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	isAdmin             = "isAdmin"
)

func (h *Handler) roling(c *gin.Context) {
	login := viper.GetString("admin.login")
	password := viper.GetString("admin.password")
	now_login, _ := c.Get("login")
	now_password, _ := c.Get("password")
	if login == now_login.(string) && password == now_password.(string) {
		c.Set(isAdmin, true)
	} else {
		c.Set(isAdmin, false)
	}
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "boş auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "tüzsüz auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "qollawçunu id'si tabılmadı")
		return 0, errors.New("qollawçunu id'si tabılmadı")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "qollawçunu id'si tüzsüz örnek bolup")
		return 0, errors.New("qollawçunu id'si tüzsüz örnek bolup")
	}

	return idInt, nil
}

func checkRole(c *gin.Context) bool {
	flag, _ := c.Get(isAdmin)
	if flag.(bool) {
		return true
	} else {
		return false
	}
}
