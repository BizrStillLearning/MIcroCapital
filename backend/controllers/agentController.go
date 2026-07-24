package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgentTrxInput struct {
	UserID uint    `json:"user_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

func GetUnverifiedUsers(c *gin.Context) {
	userRole, exists := c.Get("userRole")
	if !exists || (userRole != "agent" && userRole != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Fitur khusus Agen."})
		return
	}

	var users []models.User
	if err := config.DB.Where("role = ? AND is_verified = ?", "member", false).Order("created_at asc").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data warga"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func VerifyUser(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "agent" && userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Fitur khusus Agen."})
		return
	}

	targetID := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, targetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data warga tidak ditemukan"})
		return
	}

	user.IsVerified = true
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memverifikasi warga"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Warga atas nama " + user.Name + " berhasil diverifikasi!",
	})
}

func SearchUser(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nomor telepon wajib diisi"})
		return
	}

	var user models.User
	if err := config.DB.Where("phone = ? AND role = ?", phone, "member").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warga tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func AgentCashIn(c *gin.Context) {
	var input AgentTrxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	tx := config.DB.Begin()
	var user models.User
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, input.UserID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Warga tidak ditemukan"})
		return
	}

	user.Balance += input.Amount
	tx.Save(&user)

	transaction := models.Transaction{
		UserID:      user.ID,
		Type:        "topup",
		Amount:      input.Amount,
		Description: "Cash-In via Agen",
	}
	tx.Create(&transaction)

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Isi saldo warga berhasil", "new_balance": user.Balance})
}

func AgentCashOut(c *gin.Context) {
	var input AgentTrxInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	tx := config.DB.Begin()
	var user models.User
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, input.UserID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Warga tidak ditemukan"})
		return
	}

	if user.Balance < input.Amount {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo digital warga tidak mencukupi"})
		return
	}

	user.Balance -= input.Amount
	tx.Save(&user)

	transaction := models.Transaction{
		UserID:      user.ID,
		Type:        "withdraw",
		Amount:      input.Amount,
		Description: "Cash-Out via Agen",
	}
	tx.Create(&transaction)

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Tarik tunai warga berhasil", "new_balance": user.Balance})
}
