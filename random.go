package random

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	rd := rand.Int63()
	rs := fmt.Sprintf("Rand Fix-Print:%d", rd)
	return rs
}
