package random

import (
	"log"
	"math/rand"
	"time"
)

func Heartbeat(minMs, maxMs int) time.Duration {
	if minMs >= maxMs {
		log.Panicf("Can't generate hearbeat if minMs %d is not lower than maxMx %d", minMs, maxMs)
	}
	timeout := rand.Intn(maxMs-minMs) + minMs
	return time.Millisecond * time.Duration(timeout)
}
