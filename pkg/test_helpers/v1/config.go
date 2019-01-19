package test_helpers

import (
	"math/rand"
	"testing/quick"
	"time"
)

var DefaultConfig = &quick.Config{MaxCount: 100000}

func RandomDate(rand *rand.Rand) time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
