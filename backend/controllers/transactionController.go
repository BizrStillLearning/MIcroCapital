package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopUpInput struct {
	Phone  string  `json:"phone" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type FundInput struct {
	CampaignID uint    `json:"campaign_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required,gt=0"`
}

func TopUpBalance(c *gin.Context) {
	var input TopUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	tx := config.DB.Begin()

	var user models.User
	if err := tx.Where("phone = ?", input.Phone).First(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	user.Balance += input.Amount
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui saldo"})
		return
	}

	transaction := models.Transaction{
		UserID:      user.ID,
		Type:        "topup",
		Amount:      input.Amount,
		Description: "Isi saldo melalui Agen",
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat transaksi"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message":     "Top-Up berhasil",
		"new_balance": user.Balance,
	})
}

func FundCampaign(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := uint(userID.(float64))

	var input FundInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	tx := config.DB.Begin()

	var user models.User
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, uid).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat data pengguna"})
		return
	}

	if user.Balance < input.Amount {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo tidak mencukupi"})
		return
	}

	var campaign models.Campaign
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&campaign, input.CampaignID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Kampanye tidak ditemukan"})
		return
	}

	if campaign.Status != "active" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kampanye ini tidak menerima pendanaan"})
		return
	}

	user.Balance -= input.Amount
	campaign.CurrentAmount += input.Amount

	if campaign.CurrentAmount >= campaign.TargetAmount {
		campaign.Status = "funded"
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memotong saldo"})
		return
	}
	if err := tx.Save(&campaign).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui kampanye"})
		return
	}

	transaction := models.Transaction{
		UserID:      user.ID,
		Type:        "fund_campaign",
		Amount:      input.Amount,
		ReferenceID: campaign.ID,
		Description: "Pendanaan untuk: " + campaign.Title,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat transaksi"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message":           "Pendanaan berhasil",
		"remaining_balance": user.Balance,
		"campaign_progress": campaign.CurrentAmount,
	})
}
