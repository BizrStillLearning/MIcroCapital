package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlatformAnalytics(c *gin.Context) {
	userRole, _ := c.Get("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Otoritas Super Admin diperlukan."})
		return
	}

	var totalUsers int64
	var totalAgents int64
	var totalCampaigns int64
	var totalFundedAmount float64

	config.DB.Model(&models.User{}).Where("role = ?", "member").Count(&totalUsers)
	config.DB.Model(&models.User{}).Where("role = ?", "agent").Count(&totalAgents)
	config.DB.Model(&models.Campaign{}).Count(&totalCampaigns)

	config.DB.Model(&models.Campaign{}).Select("COALESCE(SUM(current_amount), 0)").Scan(&totalFundedAmount)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"total_users":         totalUsers,
			"total_agents":        totalAgents,
			"total_campaigns":     totalCampaigns,
			"total_funded_amount": totalFundedAmount,
		},
	})
}
