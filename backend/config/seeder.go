package config

import (
	"backend/models"
	"log"
)

func SeedSuperAdmin() {
	var admin models.User

	result := DB.Where("role = ?", "admin").First(&admin)

	if result.Error != nil {
		admin = models.User{
			Name:       "Superadmin",
			Email:      "superadmin",
			Password:   "superadmin",
			Role:       "admin",
			IsVerified: true,
		}

		DB.Create(&admin)
		log.Println("Seeder: Akun Super Admin berhasil dibuat! (Email: admin@umoja.id)")
	} else {
		log.Println("Seeder: Super Admin sudah ada. Melewati seeding.")
	}
}
