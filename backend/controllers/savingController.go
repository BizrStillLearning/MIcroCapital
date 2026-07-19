package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaySavingsFee(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := uint(userID.(float64))
	groupID := c.Param("id")

	tx := config.DB.Begin()

	var user models.User
	var group models.SavingsGroup

	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, uid).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat pengguna"})
		return
	}

	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&group, groupID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelompok tidak ditemukan"})
		return
	}

	if user.Balance < group.MonthlyFee {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo tidak mencukupi untuk bayar iuran"})
		return
	}

	user.Balance -= group.MonthlyFee
	group.TotalPool += group.MonthlyFee

	tx.Save(&user)
	tx.Save(&group)

	tx.Create(&models.Transaction{UserID: user.ID, Type: "withdraw", Amount: group.MonthlyFee, Description: "Iuran Arisan: " + group.Name})

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Iuran berhasil dibayar", "group_pool": group.TotalPool})
}
