package main

import (
	"time"

	"github.com/ZeusyBoy98/SnackYou/state"
)

func HasLock(user string) bool {
	if time.Now().After(state.Lock.Expires) {
		state.Lock.Owner = ""
		return false
	}

	return state.Lock.Owner == user
}
