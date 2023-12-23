package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet_low = "abcdefghijklmnopqrstuvwxyz"
const alphabet_up = "ABSCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = "~!@#$%^&*()-_=+;:[]{}<>?"

const all = alphabet_low + alphabet_up + numbers + symbols


func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}


func RandomString(n int) string {
	var sb strings.Builder
	k := len(all)

	for i := 0; i < n; i++ {
		c := all[rand.Intn((k))]
		sb.WriteByte(c)
	}

	return sb.String()
}