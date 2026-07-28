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

	users = idRoutes(users)
	users.GET("", userController.GetUserByID)
	users.PUT("", userController.UpdateUser)
	users.DELETE("", userController.DeleteUser)
}
func idRoutes(group *gin.RouterGroup) *gin.RouterGroup {
	return group.Group("/:id")
}
