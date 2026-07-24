package main

import (
	"backend/config"
	"backend/middlewares"
	"backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Berkas .env tidak ditemukan. Pastikan variabel sistem sudah diatur.")
	}

	config.ConnectDatabase()
	config.SeedSuperAdmin()

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	routes.SetupRoutes(r)

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Umoja API berjalan dengan baik!",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di http://localhost:%s\n", port)
	r.Run(":" + port)
}
