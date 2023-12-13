package ping

import (
	"time"
)

func (p Ping) Ping() string {
	return "pong at " + time.Now().String()
}
