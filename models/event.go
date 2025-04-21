package models

import "time"
import "restapi/db"
//import "errors"
import "fmt"


type EventStruct struct {
	ID 				int64
	Name 			string      `binding: "reguired"`
	Description 	string		`binding: "reguired"`
	Location 		string		`binding: "reguired"`
	DateTime 		time.Time	`binding: "reguired"`
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