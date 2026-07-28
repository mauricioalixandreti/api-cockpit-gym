package routes

import (
	"github.com/gin-gonic/gin"

	gymController "API-COCKPIT-GYM/Controller/gymy"
)

func setupGymRoutes(api *gin.RouterGroup) {
	gyms := api.Group("/gyms")
	gyms.POST("", gymController.CreateGym)
	gyms.GET("", gymController.GetAllGyms)

	gyms = gyms.Group("/:id")
	gyms.GET("", gymController.GetGymByID)
	gyms.PUT("", gymController.UpdateGym)
	gyms.DELETE("", gymController.DeleteGym)
}
