package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/campaigns", controllers.GetAllCampaigns)

		protected := api.Group("/")
		protected.Use(middlewares.RequireAuth())
		{
			protected.GET("/profile", func(c *gin.Context) {
				userID, _ := c.Get("userID")
				userRole, _ := c.Get("userRole")
				c.JSON(200, gin.H{"message": "Area terproteksi", "user_id": userID, "role": userRole})
			})

			protected.POST("/campaigns", controllers.CreateCampaign)
			protected.POST("/topup", controllers.TopUpBalance)
			protected.POST("/fund", controllers.FundCampaign)
			protected.POST("/loans", controllers.ApplyLoan)
			protected.POST("/loans/:id/pay", controllers.PayInstallment)
			protected.POST("/savings/:id/pay", controllers.PaySavingsFee)
			protected.GET("/agent/unverified-users", controllers.GetUnverifiedUsers)
			protected.POST("/agent/verify/:id", controllers.VerifyUser)
			protected.GET("/admin/analytics", controllers.GetPlatformAnalytics)
			protected.PUT("/profile/pin", controllers.UpdatePin)
		}
	}
}
