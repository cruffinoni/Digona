package skeleton

import (
	"math/rand"
	"time"
)

func GenerateRandomMessageColor() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%255 | rand.Int()%255<<8 | rand.Int()%255<<16
}
