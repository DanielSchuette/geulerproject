// Problem 3:
// The prime factors of 13195 are 5, 7, 13 and 29. What is
// the largest prime factor of the number 600851475143?
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.LargestPrimeFactor(600851475143)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
