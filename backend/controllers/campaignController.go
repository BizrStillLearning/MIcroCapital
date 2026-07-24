package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCampaignInput struct {
	Title        string  `json:"title" binding:"required"`
	TargetAmount float64 `json:"target_amount" binding:"required,gt=0"`
	Description  string  `json:"description" binding:"required"`
	DurationDays int     `json:"duration_days" binding:"required,gt=0"`
}

func GetAllCampaigns(c *gin.Context) {
	var campaigns []models.Campaign

	if err := config.DB.Order("created_at desc").Find(&campaigns).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kampanye"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": campaigns,
	})
}

func CreateCampaign(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak valid"})
		return
	}
	uid := uint(userID.(float64))

	var input CreateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data form tidak lengkap atau tidak valid"})
		return
	}

	campaign := models.Campaign{
		UserID:        uid,
		Title:         input.Title,
		Description:   input.Description,
		TargetAmount:  input.TargetAmount,
		CurrentAmount: 0,
		DurationDays:  input.DurationDays,
		Status:        "active",
	}

	if err := config.DB.Create(&campaign).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan kampanye ke pangkalan data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Kampanye berhasil dibuat!",
		"data":    campaign,
	})
}
