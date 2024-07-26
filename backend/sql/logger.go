package sql

import "time"

type Logger struct {
	ActionID       int64  `json:"action_id" db:"action_id"`
	User           string `db:"user" json:"user"`
	ActionEndPoint string `db:"actionEndPoint" json:"actionEndPoint"`
	Time           time.Time
}
