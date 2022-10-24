package main

import (
	"github.com/Ksis123/sa-65-example/controller"
	"github.com/Ksis123/sa-65-example/entity"
	"github.com/Ksis123/sa-65-example/middlewares"

	"github.com/gin-gonic/gin"
)

const PORT = "5000"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// Admin Routes
			router.GET("/admin", controller.ListAdmin)
			router.GET("/admin/:id", controller.GetAdmin)
			router.PATCH("/admin", controller.UpdateAdmin)
			router.DELETE("/admin/:id", controller.DeleteAdmin)

			// Banking Routes
			router.GET("/bankings", controller.ListBanking)
			router.GET("/banking/:id", controller.GetBanking)
			router.POST("/bankings", controller.CreateBanking)
			router.PATCH("/bankings", controller.UpdateBanking)
			router.DELETE("/bankings/:id", controller.DeleteBanking)

			// PaymentStatus Routes
			router.GET("/paymentstatus", controller.ListMyPaymentStatus)
			router.GET("/paymentstatus/:id", controller.GetPaymentStatus)
			router.POST("/paymentstatus", controller.CreatePaymentStatus)
			router.PATCH("/paymentstatus", controller.UpdatePaymentStatus)
			router.DELETE("/paymentstatus/:id", controller.DeletePaymentStatus)

			// Sliplist Routes
			router.GET("/Sliplist", controller.SlipList)
			router.GET("/Sliplist/:id", controller.GetSlipList)
			router.POST("/Sliplist", controller.CreateSlipList)
			router.PATCH("/Sliplist", controller.UpdateSlipList)
			router.DELETE("/Sliplist/:id", controller.DeleteSlipList)

			// StudentList Routes
			router.GET("/studentlist", controller.ListStatuses)
			router.GET("/studentlist/:id", controller.GetStudentList)
			router.POST("/studentlist", controller.CreateStudentList)
			router.PATCH("/studentlist", controller.UpdateStudentList)
			router.DELETE("/studentlist/:id", controller.DeleteStudentList)

			// Report Routes
			router.GET("/reports", controller.ListReports)
			router.GET("/reports/:id", controller.GetReport)
			router.POST("/reports", controller.CreateReport)
			router.PATCH("/reports", controller.UpdateReport)
			router.DELETE("/reports/:id", controller.DeleteReport)

			// Status Routes
			router.GET("/statuses", controller.ListStatuses)
			router.GET("/statuses/:id", controller.GetStatus)
			router.POST("/statuses", controller.CreateStatus)
			router.PATCH("/statuses", controller.UpdateReport)
			router.DELETE("/statuses/:id", controller.DeleteStatus)

			// ScholarshipStatus Routes
			router.GET("/scholarships_statuses", controller.ListMyscholarshipsStatus)
			router.GET("/scholarships_status/:id", controller.GetScholarshipsStatus)
			router.POST("/scholarships_statuses", controller.CreateScholarshipsStatus)
			router.PATCH("/scholarships_statuses", controller.UpdateScholarshipsStatus)
			router.DELETE("/scholarships_statuses/:id", controller.DeleteScholarshipsStatus)

			// ScholarshipsType Routes
			router.GET("/scholarships_types", controller.ListMyScholarshipsType)
			router.GET("/scholarships_type/:id", controller.GetScholarshipsType)
			router.POST("/scholarships_types", controller.CreateScholarshipsType)
			router.PATCH("/scholarships_types", controller.UpdateScholarshipsType)
			router.DELETE("/scholarships_types/:id", controller.DeleteScholarshipsType)

			// Scholarships Routes
			router.GET("/scholarships", controller.ListScholarships)
			router.GET("/scholarships/:id", controller.GetScholarship)
			router.POST("/scholarships", controller.CreateScholarship)
			router.PATCH("/scholarships", controller.UpdateScholarship)
			router.DELETE("/scholarships/:id", controller.DeleteScholarship)

			//Organization Router
			router.GET("/Organizations", controller.ListOrganization)
			router.GET("/Organization/:id", controller.GetOrganization)
			router.POST("/Organizations", controller.CreateOrganization)
			router.PATCH("/Organizations", controller.UpdateOrganization)
			router.DELETE("/Organization/:id", controller.DeleteOrganization)

			//TypeFund Router
			router.GET("/TypeFunds", controller.ListTypeFund)
			router.GET("/TypeFund/:id", controller.GetTypeFund)
			router.POST("/TypeFunds", controller.CreateTypeFund)
			router.PATCH("/TypeFunds", controller.UpdateTypeFund)
			router.DELETE("/TypeFund/:id", controller.DeleteTypeFund)

			//Donator Router
			router.GET("/donators", controller.ListDonators)
			router.GET("/donator/:id", controller.GetDonator)
			router.POST("/donators", controller.CreateDonator)
			router.PATCH("/donators", controller.UpdateDonator)
			router.DELETE("/donator/:id", controller.CreateDonator)

		}
	}

	// Signup User Route
	r.POST("/signup", controller.CreateAdmins)
	// // login User Route
	r.POST("/login", controller.Login)
	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
