package models

import "time"
import "restapi/db"
//import "errors"


type EventStruct struct {
	ID 				int64
	Name 			string      `binding: "reguired"`
	Description 	string		`binding: "reguired"`
	Location 		string		`binding: "reguired"`
	DateTime 		time.Time	`binding: "reguired"`
	UserID 			int			
}

var events = []EventStruct{}

func(e *EventStruct) Save() error{
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?, ?, ?, ?, ?)`
	statement, err :=  db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	//events = append(events, e)
	return nil
}

func GetAllEvents() ([]EventStruct, error) {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	var events []EventStruct
	for rows.Next(){
		var e EventStruct
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil{
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}