package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"restapi/models"
	"restapi/db"
)
func main(){
	db.InitDB()
	
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context){
	var event models.EventStruct
	err := context.ShouldBindJSON(&event)//json to struct

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
	event.Save()
}