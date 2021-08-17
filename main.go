package main

import (
	"github.com/bits2k8s/BT1_Services/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())
	r.GET("/picobrains/tokers", controllers.Tokers)
	r.GET("/picobrains/tokeme", controllers.TokeMe)
	r.POST("/picobrains", controllers.AddStdoutLine)
	r.GET("/picobrains/:toke", controllers.GetStdoutBuffer)

	r.Run(":10607")
}
