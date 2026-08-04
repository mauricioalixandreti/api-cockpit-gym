package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	Models "API-COCKPIT-GYM/Models"
	service "API-COCKPIT-GYM/Service/user"
)

func bindUserPayload(c *gin.Context, payload interface{}) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	if len(bytes.TrimSpace(body)) == 0 {
		return nil
	}

	if err := json.Unmarshal(body, payload); err == nil {
		return nil
	}

	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	if err := c.ShouldBind(payload); err == nil {
		return nil
	}

	return fmt.Errorf("invalid payload")
}

func CreateUser(c *gin.Context) {
	var payload Models.User
	if err := bindUserPayload(c, &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid payload", "error": err.Error()})
		return
	}

	user, err := service.CreateUser(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "User created successfully", "data": user})
}

func GetAllUsers(c *gin.Context) {
	users, err := service.FindAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	responseData := buildUserListResponse(users)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Users retrieved successfully", "data": responseData})
}

func GetAllUsersJSON(c *gin.Context) {
	users, err := service.FindAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, buildUserListResponse(users))
}

func buildUserListResponse(users map[string]Models.User) []Models.User {
	responseData := make([]Models.User, 0, len(users))
	for _, user := range users {
		responseData = append(responseData, user)
	}
	return responseData
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := service.FindUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User retrieved successfully", "data": user})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var payload Models.User
	if err := bindUserPayload(c, &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid payload", "error": err.Error()})
		return
	}
	payload.ID = id

	user, err := service.UpdateUser(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User updated successfully", "data": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User deleted successfully"})
}
