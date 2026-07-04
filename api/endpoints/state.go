package endpoints

import (
	"net/http"

	"github.com/ZeusyBoy98/SnackYou/state"
	"github.com/gin-gonic/gin"
)

type MoveRequest struct {
	User  string `json:"user"`
	Yaw   int    `json:"yaw"`
	Pitch int    `json:"pitch"`
}

func GetState(c *gin.Context) {
	state.Mu.RLock()
	defer state.Mu.RUnlock()

	c.JSON(http.StatusOK, gin.H{
		"turret": state.Turret,
		"lock":   state.Lock,
	})
}

func PutState(c *gin.Context) {
	var req MoveRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	state.Mu.Lock()
	defer state.Mu.Unlock()

	// lock check
	if state.Lock.Owner == "" || state.Lock.Owner != req.User {
		c.JSON(http.StatusForbidden, gin.H{"error": "no lock"})
		return
	}

	// update state
	state.Turret.Yaw = req.Yaw
	state.Turret.Pitch = req.Pitch
	state.Turret.Version++

	c.JSON(http.StatusOK, state.Turret)
}
