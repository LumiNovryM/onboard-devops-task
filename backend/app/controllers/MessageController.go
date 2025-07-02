package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"backend/app/model"
	"backend/connection"
)

func GetAllMessages(c *gin.Context) {
	var messages []model.Message

	if err := connection.DB.Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch messages",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}

func CreateMessage(c *gin.Context) {
	var input model.Message

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	if err := connection.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create message",
		})
		return
	}

	status := model.Status{
		MessageID: input.ID,
		Status:    "sent",
		UpdatedAt: time.Now(),
	}

	if err := connection.DB.Create(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Message created but failed to create status",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Message & status created successfully",
		"data": gin.H{
			"message": input,
			"status":  status,
		},
	})
}
