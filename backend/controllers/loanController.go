package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ApplyLoanInput struct {
	Title              string  `json:"title" binding:"required"`
	TotalAmount        float64 `json:"total_amount" binding:"required,gt=0"`
	MonthlyInstallment float64 `json:"monthly_installment" binding:"required,gt=0"`
}

func ApplyLoan(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := uint(userID.(float64))

	var input ApplyLoanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data pengajuan tidak valid"})
		return
	}

	loan := models.Loan{
		UserID:             uid,
		Title:              input.Title,
		TotalAmount:        input.TotalAmount,
		MonthlyInstallment: input.MonthlyInstallment,
		Status:             "pending",
		DueDate:            time.Now().AddDate(0, 1, 0),
	}

	if err := config.DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengajukan pinjaman"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pinjaman diajukan, menunggu persetujuan", "data": loan})
}

func PayInstallment(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := uint(userID.(float64))

	loanID := c.Param("id")

	tx := config.DB.Begin()

	var user models.User
	var loan models.Loan

	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&user, uid).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat data pengguna"})
		return
	}

	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("id = ? AND user_id = ?", loanID, uid).First(&loan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Pinjaman tidak ditemukan"})
		return
	}

	if loan.Status != "active" {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pinjaman ini tidak dalam status aktif"})
		return
	}

	if user.Balance < loan.MonthlyInstallment {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo tidak mencukupi untuk bayar cicilan"})
		return
	}

	user.Balance -= loan.MonthlyInstallment
	loan.PaidAmount += loan.MonthlyInstallment
	loan.DueDate = loan.DueDate.AddDate(0, 1, 0)

	if loan.PaidAmount >= loan.TotalAmount {
		loan.Status = "completed"
	}

	tx.Save(&user)
	tx.Save(&loan)

	tx.Create(&models.Transaction{UserID: user.ID, Type: "withdraw", Amount: loan.MonthlyInstallment, Description: "Bayar cicilan: " + loan.Title})

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Cicilan berhasil dibayar", "remaining_balance": user.Balance})
}
