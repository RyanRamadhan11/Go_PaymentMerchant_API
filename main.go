package main

import (
	"github.com/gin-gonic/gin"

	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/controllers"
	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/initializers"
	"github.com/RyanRamadhan11/Go_PaymentMerchant_API/middleware"
)

func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main()  {
	r := gin.Default()

	r.GET("/tes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "tes api",
		})
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.GET("/validate-login", middleware.RequireAuth, controllers.Validate)

	r.GET("/api/merchant", middleware.RequireAuth, controllers.Index)
	r.GET("/api/merchant/:id", middleware.RequireAuth, controllers.Show)
	r.POST("/api/merchant", middleware.RequireAuth, controllers.Create)
	r.PUT("/api/merchant/:id", middleware.RequireAuth, controllers.Update)
	r.DELETE("/api/merchant/:id", middleware.RequireAuth, controllers.Delete)

	r.POST("/api/payment", middleware.RequireAuth, controllers.Payment)

	r.GET("/api/user", middleware.RequireAuth, controllers.GetUsers)
	r.POST("/api/user", middleware.RequireAuth, controllers.CreateUser)
	r.GET("/api/user/:id", middleware.RequireAuth, controllers.GetUserByID)
	r.PUT("/api/user/:id", middleware.RequireAuth, controllers.UpdateUserByID)
	r.DELETE("/api/user/:id", middleware.RequireAuth, controllers.DeleteUserByID)

	r.GET("/api/history", middleware.RequireAuth, controllers.GetHistories) // Hanya ditambahkan sekali

	r.Run()
}