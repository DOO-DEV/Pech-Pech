package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	var res = make([]int, 2)
	for i := 0; ; i++ {
		if c := <-requestChan; c {
			if i == 0 || i == 1 {
				resultChan <- i + 1
				res[i] = i + 1
				continue
			}
			res = append(res, res[i-1]+res[i-2]%int(math.Pow10(9)))
			resultChan <- res[len(res)-1]
			json.Marshaler()
		}
	}

}

func main() {
	resultChan := make(chan int)
	requestChan := make(chan bool)

	go ModuloFibonacciSequence(requestChan, resultChan)
	for i := int32(0); i < 0+15; i++ {
		//start := time.Now().UnixNano()
		requestChan <- true
		new := <-resultChan
		if i < 0 {
			continue
		}
		//end := time.Now().UnixNano()
		//timeDiff := (end - start) / 1000000
		//if timeDiff < 3 {
		//	fmt.Println("Rate is too high")
		//	break
		//}
		fmt.Println(new)
	}
}
