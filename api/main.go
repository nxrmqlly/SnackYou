package main

import (
	"github.com/ZeusyBoy98/SnackYou/endpoints"
	"github.com/ZeusyBoy98/SnackYou/endpoints/esp_endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
	})

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	/*
		web:
		GET    /api/state       dashboard
		PUT    /api/state       move turret
		
		POST   /api/fire        fire snack

		GET    /api/lock        lock info
		POST   /api/lock        take control
		DELETE /api/lock        release control

		esp:
		GET    /api/esp/state   poll turret
		POST   /api/esp/ack     confirm fire
	*/

	r.GET("/api/state", endpoints.GetState)
	r.PUT("/api/state", endpoints.PutState)

	r.POST("/api/fire", endpoints.PostFire)

	r.POST("/api/lock", endpoints.PostRequestLock)
	r.DELETE("/api/lock", endpoints.DeleteReleaseLock)

	r.GET("/api/esp/state", esp_endpoints.GetESPState)
	r.POST("/api/esp/ack", esp_endpoints.PostAck)

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	r.Run()
}
