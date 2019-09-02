package main

import (
	"github.com/gin-gonic/gin"

	"my/controllers"
)

func main() {
	router := gin.Default()
	//router.Static("/assets/css", "./assets/css")
	//router.Static("/assets/javascripts", "./assets/javascripts")
	router.LoadHTMLGlob("views/*.html")

	// Settings for cookie
	// store := sessions.NewCookieStore([]byte("tmp_secret_key"))
	// router.Use(sessions.Sessions("GolangWebappTest", store))

	// Root
	router.GET("/", controllers.Index)

	router.Run(":8080")
}
