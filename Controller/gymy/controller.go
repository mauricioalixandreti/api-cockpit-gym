package gymy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	Models "API-COCKPIT-GYM/Models"

	service "API-COCKPIT-GYM/Service/gymy"
)

func CreateGym(c *gin.Context) {
	fmt.Println("Debug gym")
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

	gym, err := service.UpdateGym(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if payload.ID != c.Param("gym.ID") {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID mismatch",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Gym updated successfully",
		"data":    gym,
	})
}

func DeleteGym(c *gin.Context) {

	id := c.Param("id")
	if err := service.DeleteGym(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Gym deleted successfully",
	})
}

func GetGymByOwnerEmail(c *gin.Context) {

	email := c.Query("owner_email")

	// Validar se o email foi informado.
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{

			"sucess":  false,
			"message": "O email do proprietário não foi informado",
		})
		return
	}

	gymData, err := service.FindGymByOwnerEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{

			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Retornar a academia encontrada.
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Academia encontrada com sucesso",
		"data":    gymData,
	})
}
