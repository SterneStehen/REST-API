package main

import(
	//"net/http"
	"github.com/gin-gonic/gin"
	//"restapi/models"
	"restapi/routes"
	"restapi/db"
	//"strconv"
)
func main(){
	db.InitDB()
	
	server := gin.Default()
	defer db.DB.Close()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
