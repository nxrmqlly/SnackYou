package esp_endpoints

import (
	"net/http"

	"github.com/ZeusyBoy98/SnackYou/state"
	"github.com/gin-gonic/gin"
)

func GetESPState(c *gin.Context) {
	state.Mu.RLock()
	defer state.Mu.RUnlock()

	//? ESP only needs current command
	c.JSON(http.StatusOK, gin.H{
		"yaw":     state.Turret.Yaw,
		"pitch":   state.Turret.Pitch,
		"fire":    state.Turret.Fire,
		"version": state.Turret.Version,
	})
}
