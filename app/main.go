package main

import (
	"github.com/gin-gonic/gin"

	"my/controllers"
)

func main() {
	router := gin.Default()
	router.Static("/assets/stylesheets", "./assets/stylesheets")
	router.Static("/assets/images", "./assets/images")
	//router.Static("/assets/javascripts", "./assets/javascripts")
	router.LoadHTMLGlob("views/*.html")

	// Settings for cookie
	// store := sessions.NewCookieStore([]byte("tmp_secret_key"))
	// router.Use(sessions.Sessions("GolangWebappTest", store))

	// Root
	router.GET("/", controllers.Index)
	router.GET("/reset", controllers.Reset)
	router.POST("/execute", controllers.Execute)

	router.Run(":8080")
}
