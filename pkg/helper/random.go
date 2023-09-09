package helper

import (
	"math/rand"
	"time"
)

const stringBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const numberBytes = "1234567890"

func RandomNumber(n int) string {

	if n <= 0 {
		return ""
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}

	return string(b)
}

func RandomChain(n int) string {
	if n <= 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)

	for i := range b {
		b[i] = stringBytes[rand.Intn(len(stringBytes))]
	}

	return string(b)
}
