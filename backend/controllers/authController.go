package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	PIN   string `json:"pin" binding:"required,min=4,max=6"`
}

type LoginInput struct {
	Phone string `json:"phone" binding:"required"`
	PIN   string `json:"pin" binding:"required"`
}

type UpdatePinInput struct {
	CurrentPin string `json:"current_pin" binding:"required"`
	NewPin     string `json:"new_pin" binding:"required,min=4"`
}

func UpdatePin(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid := uint(userID.(float64))

	var input UpdatePinInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if user.PIN != input.CurrentPin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "PIN saat ini tidak cocok"})
		return
	}

	user.PIN = input.NewPin
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan PIN baru"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PIN berhasil diperbarui"})
}

type AdminLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func AdminLogin(c *gin.Context) {
	var input AdminLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format email atau password salah"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak terdaftar"})
		return
	}

	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Anda bukan Admin."})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token server"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Admin berhasil",
		"token":   token,
		"user":    user,
	})
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak lengkap atau format salah"})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("phone = ?", input.Phone).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Nomor telepon sudah terdaftar"})
		return
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(input.PIN), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses PIN"})
		return
	}

	user := models.User{
		Name:  input.Name,
		Phone: input.Phone,
		PIN:   string(hashedPin),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendaftarkan pengguna"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil! Silakan kunjungi Agen Lokal untuk verifikasi akun.",
	})
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nomor HP dan PIN wajib diisi"})
		return
	}

	var user models.User
	if err := config.DB.Where("phone = ?", input.Phone).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nomor telepon atau PIN salah"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PIN), []byte(input.PIN)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nomor telepon atau PIN salah"})
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "rahasia_umoja_super_aman"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token autentikasi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   tokenString,
		"user": gin.H{
			"id":          user.ID,
			"name":        user.Name,
			"role":        user.Role,
			"is_verified": user.IsVerified,
		},
	})
}
