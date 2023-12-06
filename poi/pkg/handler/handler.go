package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"restapi/pkg/service"
	"restapi/pkg/service/wb"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(wsHandler *wb.Handler) *gin.Engine {
	router := gin.New()
	adminHandler := router.Group("/admin/user")
	{

		adminHandler.GET("/all", h.GetUsers)
		adminHandler.POST("/", h.CreateUser)
		adminHandler.GET("/", h.GetUserByEmail)
	}

	userHandler := router.Group("/user")
	{
		userHandler.POST("/register", h.Register)
		userHandler.POST("/login", h.Login)
	}

	apiHandler := router.Group("/api", h.userIdentity)
	{
		imageHandler := apiHandler.Group("/images")
		{
			imageHandler.POST("/", h.createAva)
			imageHandler.GET("/", h.getAllImages)
			imageHandler.GET("/:id", h.getImageById)
			imageHandler.DELETE("/:id", h.deleteImage)
			imageHandler.GET("/metrics", gin.WrapH(promhttp.Handler()))
		}

		chatHandler := apiHandler.Group("/chat")
		{
			chatHandler.POST("/ws/createRoom", wsHandler.CreateRoom)
			chatHandler.GET("/ws/getRooms", wsHandler.GetRooms)
			chatHandler.GET("/ws/getClients/:roomId", wsHandler.GetClients)
		}
	}
	router.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	return router
}
