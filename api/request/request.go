package request

type MoveRequest struct {
	User  string `json:"user"`
	Yaw   int    `json:"yaw"`
	Pitch int    `json:"pitch"`
}

type FireRequest struct {
	User string `json:"user"`
}

type LockRequest struct {
	User string `json:"user"`
}

type AckRequest struct {
	Seq int `json:"seq"`
}
