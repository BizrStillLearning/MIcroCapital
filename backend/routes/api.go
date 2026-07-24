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
		api.POST("/admin/login", controllers.AdminLogin)

		api.GET("/campaigns", controllers.GetAllCampaigns)

		protected := api.Group("/")
		protected.Use(middlewares.RequireAuth())
		{
			protected.POST("/campaigns", controllers.CreateCampaign)
			protected.POST("/topup", controllers.TopUpBalance)
			protected.POST("/fund", controllers.FundCampaign)
			protected.POST("/loans", controllers.ApplyLoan)
			protected.GET("/loans", controllers.GetMyLoans)
			protected.POST("/loans/:id/pay", controllers.PayInstallment)
			protected.POST("/savings/:id/pay", controllers.PaySavingsFee)
			protected.GET("/agent/unverified-users", controllers.GetUnverifiedUsers)
			protected.POST("/agent/verify/:id", controllers.VerifyUser)
			protected.GET("/agent/search-user", controllers.SearchUser)
			protected.POST("/agent/cash-in", controllers.AgentCashIn)
			protected.POST("/agent/cash-out", controllers.AgentCashOut)
			protected.GET("/admin/analytics", controllers.GetPlatformAnalytics)
			protected.PUT("/profile/pin", controllers.UpdatePin)
			protected.GET("/transactions/history", controllers.GetHistory)
			protected.GET("/profile", controllers.GetProfile)
			protected.POST("/pay-kas", controllers.PayIuranKas)
			protected.GET("/admin/loans/pending", controllers.GetPendingLoans)
			protected.POST("/admin/loans/:id/approve", controllers.ApproveLoan)
			protected.GET("/admin/agents", controllers.GetAgents)
			protected.POST("/admin/agents/:id/approve", controllers.ApproveAgent)
		}
	}
}
