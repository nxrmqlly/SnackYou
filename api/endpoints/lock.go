package endpoints

import (
	"net/http"
	"time"

	"github.com/ZeusyBoy98/SnackYou/request"
	"github.com/ZeusyBoy98/SnackYou/state"

	"github.com/gin-gonic/gin"
)

// POST
func PostRequestLock(c *gin.Context) {

	var req request.LockRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	state.Mu.Lock()
	defer state.Mu.Unlock()

	if time.Now().After(state.Lock.Expires) {
		state.Lock.Owner = ""
	}

	if state.Lock.Owner != "" && state.Lock.Owner != req.User {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"owner":   state.Lock.Owner,
		})
		return
	}

	state.Lock.Owner = req.User
	state.Lock.Expires = time.Now().Add(30 * time.Second)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"expires": state.Lock.Expires,
	})
}

// DELETE
func DeleteReleaseLock(c *gin.Context) {
	user := c.Query("user")

	state.Mu.Lock()
	defer state.Mu.Unlock()

	if state.Lock.Owner == user {
		state.Lock.Owner = ""
		// set expires to start of epoch so it basically is always expired
		state.Lock.Expires = time.Unix(0, 0)
	}

	c.JSON(http.StatusOK, gin.H{"released": true})
}

// GET
func GetLock(c *gin.Context) {
	state.Mu.RLock()
	defer state.Mu.RUnlock()

	c.JSON(http.StatusOK, state.Lock)
}
