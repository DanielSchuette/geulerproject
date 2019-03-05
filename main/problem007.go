// Problem 7:
// By listing the first six prime numbers: 2, 3, 5, 7, 11,
// and 13, we can see that the 6th prime is 13. What is the
// 10 001st prime number?
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.FindPrimeNumber(10001)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
