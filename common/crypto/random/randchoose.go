package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandChoose(list []rune) rune {
	return list[rand.Intn(len(list))]
}
