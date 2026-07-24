package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlatformAnalytics(c *gin.Context) {
	userRole, exists := c.Get("userRole")
	if !exists || userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Area khusus Administrator."})
		return
	}

	var totalUsers int64
	var totalAgents int64
	var totalFundedAmount float64

	config.DB.Model(&models.User{}).Where("role = ?", "member").Count(&totalUsers)

	config.DB.Model(&models.User{}).Where("role = ?", "agent").Count(&totalAgents)

	config.DB.Table("campaigns").Select("COALESCE(SUM(current_amount), 0)").Scan(&totalFundedAmount)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"total_users":         totalUsers,
			"total_agents":        totalAgents,
			"total_funded_amount": totalFundedAmount,
		},
	})
}

func GetPendingLoans(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
		return
	}

	var loans []models.Loan
	if err := config.DB.Preload("User").Where("status = ?", "pending").Order("created_at asc").Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar antrean pinjaman"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": loans})
}

func ApproveLoan(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
		return
	}

	loanID := c.Param("id")

	tx := config.DB.Begin()

	var loan models.Loan
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&loan, loanID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Pinjaman tidak ditemukan"})
		return
	}

	if loan.Status != "pending" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pinjaman ini sudah diproses sebelumnya"})
		return
	}

	var user models.User
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, loan.UserID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat data warga"})
		return
	}

	loan.Status = "active"

	user.Balance += loan.TotalAmount

	transaction := models.Transaction{
		UserID:      user.ID,
		Type:        "topup",
		Amount:      loan.TotalAmount,
		Description: "Pencairan Pinjaman: " + loan.Title,
	}

	tx.Save(&loan)
	tx.Save(&user)
	tx.Create(&transaction)

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "Pinjaman berhasil disetujui. Dana telah dicairkan ke dompet warga.",
	})
}

func GetAgents(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
		return
	}

	var agents []models.User
	if err := config.DB.Where("role = ?", "agent").Order("created_at desc").Find(&agents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data agen"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": agents})
}

func ApproveAgent(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
		return
	}

	agentID := c.Param("id")

	var agent models.User
	if err := config.DB.Where("id = ? AND role = ?", agentID, "agent").First(&agent).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data agen tidak ditemukan"})
		return
	}

	agent.IsVerified = true
	if err := config.DB.Save(&agent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyetujui agen"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Agen berhasil disetujui dan kini aktif."})
}
