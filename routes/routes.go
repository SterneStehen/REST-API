package routes

import(
	//"net/http"
	"github.com/gin-gonic/gin"
// 	"restapi/models"
// 	"restapi/db"
// 	"strconv"
 )

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
}