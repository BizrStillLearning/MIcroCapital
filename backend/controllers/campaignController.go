package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCampaignInput struct {
	Title        string  `json:"title" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	TargetAmount float64 `json:"target_amount" binding:"required,gt=0"`
	DurationDays int     `json:"duration_days" binding:"required,gt=0"`
}

func CreateCampaign(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak terautentikasi"})
		return
	}

	var input CreateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data kampanye tidak valid"})
		return
	}

	uid := uint(userID.(float64))

	campaign := models.Campaign{
		UserID:       uid,
		Title:        input.Title,
		Description:  input.Description,
		TargetAmount: input.TargetAmount,
		DurationDays: input.DurationDays,
		Status:       "pending",
	}

	if err := config.DB.Create(&campaign).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan kampanye"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Kampanye berhasil dibuat dan menunggu peninjauan",
		"data":    campaign,
	})
}

func GetAllCampaigns(c *gin.Context) {
	var campaigns []models.Campaign

	if err := config.DB.Preload("User").Where("status = ?", "active").Find(&campaigns).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kampanye"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": campaigns,
	})
}
