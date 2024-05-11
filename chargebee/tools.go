package chargebee

import (
	"math/rand"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(size int) string {
	var s string
	for i := 0; i < size; i++ {
		s += string(letters[rand.Intn(len(letters))])
	}
	return s
}
