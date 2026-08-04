package routes

import (
	"github.com/gin-gonic/gin"

	userController "API-COCKPIT-GYM/Controller/user"
)

func setupUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	users.POST("", userController.CreateUser)
	users.GET("", userController.GetAllUsers)
	users.GET("/json", userController.GetAllUsersJSON)

	users.GET("/:id", userController.GetUserByID)
	users.PUT("/:id", userController.UpdateUser)
	users.DELETE("/:id", userController.DeleteUser)
}
