// Problem 1:
// If we list all the natural numbers below 10 that
// are multiples of 3 or 5, we get 3, 5, 6 and 9. The
// sum of these multiples is 23. Find the sum of all
// the multiples of 3 or 5 below 1000.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.FindMultiples(3, 5, 0, 1000)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
