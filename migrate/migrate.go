package main

import (
	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/edelwei88/fixed-interest-go/models"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectToDB()
}

func main() {
	initialize.DB.AutoMigrate(&models.Role{}, &models.LoanType{}, &models.User{}, &models.Docs{}, &models.Loan{}, &models.LoanPayment{}, &models.Token{})
}
