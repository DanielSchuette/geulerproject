// Problem 5:
// 2520 is the smallest number that can be divided by each
// of the numbers from 1 to 10 without any remainder. What
// is the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20?
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.EvenlyDivisible(20)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
