package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"restapi/models"
	"restapi/db"
	"strconv"

)
func main(){
	db.InitDB()
	
	server := gin.Default()
	defer db.DB.Close()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
	}
	context.JSON(http.StatusOK, events) //respons in JSON
}

func getEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
	}
	context.JSON(http.StatusOK, event)
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
	//context.JSON(http.StatusBadRequest, gin.H{"message": "Event created!", "event": event})
	err = event.SaveToDb()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event. "})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}