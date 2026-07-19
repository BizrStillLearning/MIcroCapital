package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUnverifiedUsers(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "agent" && userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Hanya untuk Agen."})
		return
	}

	var users []models.User
	if err := config.DB.Where("is_verified = ? AND role = ?", false, "member").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data warga"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func VerifyUser(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "agent" && userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
		return
	}

	targetUserID := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, targetUserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	user.IsVerified = true
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memverifikasi pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warga berhasil diverifikasi", "user_id": user.ID})
}
