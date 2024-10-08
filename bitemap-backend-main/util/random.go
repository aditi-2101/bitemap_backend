package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generates random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generates random string of length n.
func RandomString(n int) *string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	randomString := sb.String()
	return &randomString
}

// generates random owner name.
func RandomOwner() *string {
	return RandomString(6)
}

// generate random email
func RandomEmail() *string {
	randEmail := fmt.Sprintf("%s@gmail.com", *(RandomString(6)))
	return &randEmail
}
