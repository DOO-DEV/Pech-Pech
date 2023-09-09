package main

import (
	"fmt"
	"math/rand"
)

func main() {

	k := make([]byte, 5)
	const numberBytes = "1234567890"
	for i := range k {
		k[i] = numberBytes[rand.Intn(len(numberBytes))]
	}
	fmt.Println(string(k))
}
