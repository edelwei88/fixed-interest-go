package main

import (
	"github.com/edelwei88/fixed-interest-go/controllers"
	"github.com/edelwei88/fixed-interest-go/initialize"
	"github.com/gin-gonic/gin"
)

func setup() {
	initialize.LoadEnv()
	initialize.ConnectToDB()
}

func main() {
	setup()
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Roles
	router.GET("/roles", controllers.RolesGET)
	router.GET("/roles/:id", controllers.RoleGET)
	router.POST("/roles", controllers.RolePOST)
	router.PATCH("/roles/:id", controllers.RolePATCH)
	router.DELETE("/roles/:id", controllers.RoleDELETE)

	// Docs
	// router.GET("/docs", controllers.DocsGET)
	// router.GET("/docs/:id", controllers.DocGET)
	// router.POST("/docs", controllers.DocsPOST)
	// router.PATCH("/docs/:id", controllers.DocsPATCH)
	// router.GET("/docs/:id", controllers.DocsDELETE)

	// admin := router.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	os.Getenv("ADMIN_USERNAME"): os.Getenv("ADMIN_PASSWORD"),
	// }))
	// admin.GET("/test", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"rofl": "rofl",
	// 	})
	// })

	router.Run()
}
