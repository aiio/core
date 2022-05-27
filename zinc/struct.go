package zinc

import "time"

type Version struct {
	Branch     string    `json:"Branch"`
	Build      string    `json:"Build"`
	BuildDate  time.Time `json:"BuildDate"`
	CommitHash string    `json:"CommitHash"`
	Version    string    `json:"Version"`
}
