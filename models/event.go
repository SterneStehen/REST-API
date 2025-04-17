package madels

import "time"


type Event struct {
	ID 				int
	Name 			string     `binding: "reguired"`
	Description 	string		`binding: "reguired"`
	Location 		string		`binding: "reguired"`
	DateTime 		time.Time	`binding: "reguired"`
	UserID 			int			`binding: "reguired"`
}

var events = []Event{}

func(e Event) Save(){
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}