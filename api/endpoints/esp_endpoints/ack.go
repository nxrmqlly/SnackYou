package esp_endpoints

import (
	"net/http"

	"github.com/ZeusyBoy98/SnackYou/state"
	"github.com/gin-gonic/gin"
)

type AckRequest struct {
	Version int `json:"version"`
}

func PostAck(c *gin.Context) {
	var req AckRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	state.Mu.Lock()
	defer state.Mu.Unlock()

	// only clear fire if matching version
	if state.Turret.Version == req.Version {
		state.Turret.Fire = false
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
