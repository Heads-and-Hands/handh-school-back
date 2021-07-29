package config

import (
	"os"
	"time"
)

var StartDelay time.Duration
var CheckTimer time.Duration
var Timeout time.Duration

func init() {
	s := os.Getenv("START_DELAY")
	StartDelay, _ = time.ParseDuration(s)
}
