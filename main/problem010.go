// Problem 10:
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.SumOfPrimes(2000000)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
