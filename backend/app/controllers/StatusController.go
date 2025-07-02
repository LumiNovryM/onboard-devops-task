package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/app/model"
	"backend/connection"
)

func GetAllStatus(c *gin.Context) {
	var status []model.Status

	if err := connection.DB.Find(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch messages",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": status,
	})
}
