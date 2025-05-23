package db

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3" 
	
	//"fmt"
	
)

var DB *sql.DB

func InitDB(){
	
	var err error
	driverName := "sqlite3"
	DB, err = sql.Open(driverName, "api.db")
	if err != nil{
		panic("Cannot connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables(){
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)
	if err != nil{
		panic("Could not created User table.")
	}


	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil{
		panic("Could not created Event table.")
	}
}

