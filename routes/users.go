package routes

	import(
	"net/http"
	"github.com/gin-gonic/gin"
	"restapi/models"
	"restapi/utils"
	//"restapi/db"
	//"strconv"
	)

func registerUser(context *gin.Context){
	var user models.UserStruct
	err := context.ShouldBindJSON(&user)//json to struct
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	user.ID = 1
	err = user.SaveUserToDb()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save User. "})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Created user tablet. "})
}

func loginUser(context *gin.Context){
	var user models.UserStruct
	err := context.ShouldBindJSON(&user)//json to struct
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	err = user.Validate()
	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not created token."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!",  "token": token})
	}	
