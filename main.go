package main

import (
	"net/http"
	"time"

	"github.com/CaiqueCastro07/portfolio-backend-go/src/controllers"
	"github.com/CaiqueCastro07/portfolio-backend-go/src/database"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.IndentedJSON(204, "") // use this return, AbortWithStatus(204) will overwrite the headers
			return
		}
		c.Next()
	}
}

func main() {

	database.InitDatabase()

	router := gin.New()
	router.Use(CORS())

	router.GET("/goapi/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, map[string]interface{}{
			"status": 200, "online": true, "integration": "Portfolio Golang", "time": time.Now()})
	})

	router.POST("/goapi/message", controllers.RegisterMessage)

	router.Run("0.0.0.0:3002") // change to localhost:3002 when running locally

}
