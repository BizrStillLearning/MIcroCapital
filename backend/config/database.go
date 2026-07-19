package config

import (
	"fmt"
	"log"
	"os"

	"backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = database
	fmt.Println("Berhasil terhubung ke database MySQL!")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Campaign{},
		&models.Transaction{},
		&models.Loan{},
		&models.SavingsGroup{},
	)
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}
	fmt.Println("Migrasi tabel selesai!")
}
