package handler

import (
	"github.com/AlibekDalgat/pos-credition/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)

	}
	api := router.Group("/api", h.userIdentity)
	{
		admin := api.Group("/admin")
		{
			shops := admin.Group("/shops")
			{
				shops.POST("/", h.createShop)
				shops.GET("/", h.getAllShops)
				shops.GET("/:id", h.getShopById)
				shops.PUT("/:id", h.updateShop)
				shops.DELETE("/:id", h.deleteShop)
			}

			market_places := admin.Group("/market_places")
			{
				market_places.POST("/", h.createMarketPlace)
				market_places.GET("/", h.getAllMarketPlaces)
				market_places.GET("/:id", h.getMarketPlaceById)
				market_places.PUT("/:id", h.updateMarketPlace)
				market_places.DELETE("/:id", h.deleteMarketPlace)
			}

			agents := admin.Group("/agents")
			{
				agents.POST("/", h.createAgent)
				agents.GET("/", h.getAllAgents)
				agents.GET("/:id", h.getAgentById)
				agents.PUT("/:id", h.updateAgent)
				agents.DELETE("/:id", h.deleteAgent)
				agents.PATCH("/:id", h.newAccess)
			}

		}

		/*
			items := api.Group("/agent")
			{
				items.PUT("/", h.createCredits)
				items.GET("/:id", h.getAllCredits)
				items.PUT("/:id", h.deleteCredit)
			}

		*/

	}
	return router
}
