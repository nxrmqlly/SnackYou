package state

import (
	"sync"
	"time"
)

type TurretState struct {
	Yaw     int  `json:"yaw"`
	Pitch   int  `json:"pitch"`
	Fire    bool `json:"fire"`
	Version int  `json:"version"`
}

type LockState struct {
	Owner   string    `json:"owner"`
	Expires time.Time `json:"expires"`
}

var (
	Mu sync.RWMutex

	Turret = TurretState{
		Yaw:   90,
		Pitch: 20,
	}

	Lock = LockState{}
)
