package routes

	import(
	"net/http"
	"github.com/gin-gonic/gin"
	"restapi/models"
	//"restapi/db"
	"strconv"
	)

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
	context.JSON(http.StatusOK, gin.H{"message": "Created event. "})
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
		return
	}
	context.JSON(http.StatusOK, event)
}




func updateEvent(context *gin.Context){

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}

	err = models.SetEventByID(event, context)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message": "Update event successfully.", "event": event})

}