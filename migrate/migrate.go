package main

import (
	"github.com/edelwei88/fixed-interest-go/initial"
	"github.com/edelwei88/fixed-interest-go/models"
)

func init() {
	initial.LoadEnv()
	initial.ConnectToDB()
}

func main() {
	initial.DB.AutoMigrate(&models.Role{})
}
