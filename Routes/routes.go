package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	setupUserRoutes(api)
	setupGymRoutes(api)
	PaymentsRoutes(api)

}
