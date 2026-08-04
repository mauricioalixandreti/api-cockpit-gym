package routes

import (
	"github.com/gin-gonic/gin"
	gymController "API-COCKPIT-GYM/Controller/gymy"
)

func setupGymRoutes(api *gin.RouterGroup) {
	gyms := api.Group("/gyms")
	gyms.POST("", gymController.CreateGym)
	gyms.GET("", gymController.GetAllGyms)

	gyms.GET("/:id", gymController.GetGymByID)
	gyms.PUT("/:id", gymController.UpdateGym)
	gyms.DELETE("/:id", gymController.DeleteGym)
}
