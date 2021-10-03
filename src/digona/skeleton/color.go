package skeleton

import "math/rand"

func GenerateRandomMessageColor() int {
	return rand.Int()%255 | rand.Int()%255<<8 | rand.Int()%255<<16
}
