package models

import "time"
import "restapi/db"
//import "errors"
import "fmt"
//import "net/http"
import	"github.com/gin-gonic/gin"


type EventStruct struct {
	ID 				int64
	Name 			string      `binding:"required"`
	Description 	string		`binding:"required"`
	Location 		string		`binding:"required"`
	DateTime 		time.Time	`binding:"required"`
	UserID 			int			
}

var events = []EventStruct{}

func(e *EventStruct) SaveToDb() error{
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?, ?, ?, ?, ?)`
	statement, err :=  db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer statement.Close()
	resultExec, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil{
		return err
	}
	id, err := resultExec.LastInsertId()
	e.ID = id
	//events = append(events, e)
	return nil
}

func GetAllEvents() ([]EventStruct, error) {
	queryResult, err := db.DB.Query("SELECT * FROM events")
	if err != nil{
		return nil, err
	}
	defer queryResult.Close()
	var events []EventStruct
	for queryResult.Next(){
		var e EventStruct
		err := queryResult.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil{
			return nil, err
		}
		fmt.Println("id = ", e.ID, "name = ", e.Name)
		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(id int64)(*EventStruct, error){
	query := "SELECT * FROM events WHERE id = ?"
	queryResult := db.DB.QueryRow(query, id)
	var e EventStruct
	err := queryResult.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil{
		return nil,	err
	}
	return &e, nil
}

 func SetEventByID(e *EventStruct, context *gin.Context) (error){
	//var e EventStruct
	err := context.ShouldBindJSON(&e)//json to struct
	if err != nil{
		//context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return err
	}
	//e.ID = eventId
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?`
	statement, err :=  db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil{
		return err
	}
	return nil
 }