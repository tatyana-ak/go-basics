package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authentificated := server.Group("/")
	authentificated.Use(middlewares.Authenticate)
	authentificated.PUT("/events/:id", updateEvent)
	authentificated.DELETE("/events/:id", deleteEvent)

	server.POST("/events", middlewares.Authenticate, createEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
