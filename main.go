package main

import (
	"github.com/edelwei88/fixed-interest-go/controllers"
	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/edelwei88/fixed-interest-go/middlewares"
	"github.com/gin-gonic/gin"
)

func setup() {
	initialize.LoadEnv()
	initialize.ConnectToDB()
}

func main() {
	setup()
	router := gin.Default()

	// Docs
	{
		docs := router.Group("/docs")
		docs.GET("/", controllers.DocsGET)
		docs.GET("/:id", controllers.DocGET)
		docs.POST("/", controllers.DocsPOST)
		docs.PATCH("/:id", controllers.DocsPATCH)
		docs.DELETE("/:id", controllers.DocsDELETE)
	}

	// Loans
	{
		loans := router.Group("/loans")
		loans.GET("/", controllers.LoansGET)
		loans.GET("/:id", controllers.LoanGET)
		loans.POST("/", controllers.LoanPOST)
		loans.PATCH("/:id", controllers.LoanPATCH)
		loans.DELETE("/:id", controllers.LoanDELETE)
	}

	// LoanPayments
	{
		loanPayments := router.Group("/loan_payments")
		loanPayments.GET("/", controllers.LoanPaymentsGET)
		loanPayments.GET("/:id", controllers.LoanPaymentGET)
		loanPayments.POST("/", controllers.LoanPaymentPOST)
		loanPayments.PATCH("/:id", controllers.LoanPaymentPATCH)
		loanPayments.DELETE("/:id", controllers.LoanPaymentDELETE)
	}

	// LoanTypes
	{
		loanTypes := router.Group("/loan_types")
		loanTypes.GET("/", controllers.LoanTypesGET)
		loanTypes.GET("/:id", controllers.LoanTypeGET)
		loanTypes.POST("/", controllers.LoanTypePOST)
		loanTypes.PATCH("/:id", controllers.LoanTypePATCH)
		loanTypes.DELETE("/:id", controllers.LoanTypeDELETE)
	}

	// Roles
	{
		roles := router.Group("/roles")
		roles.GET("/", controllers.RolesGET)
		roles.GET("/:id", controllers.RoleGET)
		roles.POST("/", controllers.RolePOST)
		roles.PATCH("/:id", controllers.RolePATCH)
		roles.DELETE("/:id", controllers.RoleDELETE)
	}

	// Users
	{
		users := router.Group("/users")
		users.GET("/", controllers.UsersGET)
		users.GET("/:id", controllers.UserGET)
		users.POST("/", controllers.UserPOST)
		users.PATCH("/:id", controllers.UserPATCH)
		users.DELETE("/:id", controllers.UserDELETE)
	}

	// Authorization
	{
		auth := router.Group("/auth")
		auth.POST("/login", controllers.LoginPOST)
		auth.POST("/register", controllers.RegisterPOST)
	}

	// User
	{
		data := router.Group("/data", middlewares.BearerTokenAuth())
		data.GET("/current_user", controllers.CheckBearerTokenGET)
		data.POST("/add_loan", controllers.AddLoanPOST)
	}

	router.Run()
}
