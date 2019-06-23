package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	port := os.Getenv("PORT")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hey!": "Message from Heroku, with restart now 🔥🔥",
		})
	})
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("error in starting server", err)
	}
}
