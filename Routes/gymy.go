package routes

import (
	gymController "API-COCKPIT-GYM/Controller/gymy"

	"github.com/gin-gonic/gin"
)

func setupGymRoutes(api *gin.RouterGroup) {
	gyms := api.Group("/gyms")
	gyms.POST("", gymController.CreateGym)
	gyms.GET("", gymController.GetAllGyms, gymController.GetGymByOwnerEmail)

	gyms.GET("/:id", gymController.GetGymByID)
	gyms.PUT("/:id", gymController.UpdateGym)
	gyms.DELETE("/:id", gymController.DeleteGym)
}
