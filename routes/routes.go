package routes

import (
	"github.com/gin-gonic/gin"
	"org.gfoo/api-rest/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticated)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

}
