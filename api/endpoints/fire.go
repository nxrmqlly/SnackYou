package endpoints

import (
	"net/http"

	"github.com/ZeusyBoy98/SnackYou/state"
	"github.com/gin-gonic/gin"
)

type FireRequest struct {
	User string `json:"user"`
}

func PostFire(c *gin.Context) {
	var req FireRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	state.Mu.Lock()
	defer state.Mu.Unlock()

	if state.Lock.Owner == "" || state.Lock.Owner != req.User {
		c.JSON(http.StatusForbidden, gin.H{"error": "no lock"})
		return
	}

	state.Turret.Fire = true
	state.Turret.Version++

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"version": state.Turret.Version,
	})
}
