package gymy

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Models "API-COCKPIT-GYM/Models"
	service "API-COCKPIT-GYM/Service/gymy"
)

func CreateGym(c *gin.Context) {
	var payload Models.Gym
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid payload"})
		return
	}

	gym, err := service.CreateGym(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Gym created successfully", "data": gym})
}

func GetAllGyms(c *gin.Context) {
	gyms, err := service.FindAllGyms(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Gyms retrieved successfully", "data": gyms})
}

func GetGymByID(c *gin.Context) {
	id := c.Param("id")
	gym, err := service.FindGymByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Gym retrieved successfully", "data": gym})
}

func UpdateGym(c *gin.Context) {
	var payload Models.Gym
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid payload"})
		return
	}
	payload.ID = c.Param("id")

	gym, err := service.UpdateGym(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Gym updated successfully", "data": gym})
}

func DeleteGym(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteGym(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Gym deleted successfully"})
}
