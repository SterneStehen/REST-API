package models

//import "time"
import "restapi/db"
import "restapi/utils"
import "errors"
import "fmt"
import "database/sql"
//import "net/http"
//import	"github.com/gin-gonic/gin"


type UserStruct struct {
	ID 				int64
	Email 			string      `binding:"required"`
	Password 		string		`binding:"required"`
	
}


func(u *UserStruct) SaveUserToDb() error{
	query := `
	INSERT INTO users(email, password) 
	VALUES(?, ?)`
	statement, err :=  db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer statement.Close()
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}
	resultExec, err := statement.Exec(u.Email, hashedPass)
	if err != nil{
		return err
	}
	id, err := resultExec.LastInsertId()
	u.ID = id
	//events = append(events, e)
	return nil
}

func (u UserStruct) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row   := db.DB.QueryRow(query, u.Email)

	var dbPassword string
	if err := row.Scan(&u.ID, &dbPassword)
	err != nil {
    	if err == sql.ErrNoRows {
        	return errors.New("user not found")
    	}
    	return err
	}

	
	passIsValid := utils.CheckPasswordHash(u.Password, dbPassword)
	fmt.Println(passIsValid)
	if !passIsValid{
		return errors.New("Credentials invalid")
	}
	return nil
}