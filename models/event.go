package models

import "time"


type EventStruct struct {
	ID 				int
	Name 			string      `binding: "reguired"`
	Description 	string		`binding: "reguired"`
	Location 		string		`binding: "reguired"`
	DateTime 		time.Time	`binding: "reguired"`
	UserID 			int			
}

var events = []EventStruct{}

func(e EventStruct) Save(){
	events = append(events, e)
}

func GetAllEvents() []EventStruct {
	return events
}