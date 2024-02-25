package routes

import (
	"backenddemo/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/getallCharacter", controllers.GetallCharacter)
	r.GET("/getcharacterbyID/:id", controllers.GetcharacterbyID)
	r.POST("/createCharacter", controllers.CreateCharacter)
	r.PUT("/updateCharacter", controllers.UpdateCharacter)
	r.DELETE("deleteCharacterbyID", controllers.DeleteCharacterbyID)
}
