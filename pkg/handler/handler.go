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
				shops.PUT("/:id", h.updateShop)
				shops.POST("/:id", h.deleteShop)
			}

			/*
				market_places := admin.Group("/market_places")
				{
					market_places.PUT("/", h.createMark_place)
					market_places.GET("/", h.getAllMark_place)
					//market_places.PUT("/:id", h.updateMark_place)
					//market_places.POST("/:id", h.deleteMark_place)
				}

				agents := admin.Group("/agents")
				{
					market_places.PUT("/", h.createAgent)
					market_places.GET("/", h.getAllAgent)
					market_places.PUT("/:id", h.updateAgent)
					market_places.POST("/:id", h.deleteAgent)
				}

			*/
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
