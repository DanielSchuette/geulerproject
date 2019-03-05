// Problem 6:
// The sum of the squares of the first ten natural numbers
// is, 1^2 + 2^2 + ... + 10^2 = 385. The square of the sum of
// the first ten natural numbers is, (1 + 2 + ... + 10)^2 =
// 552 = 3025. Hence the difference between the sum of the
// squares of the first ten natural numbers and the square of
// the sum is 3025 − 385 = 2640. Find the difference between
// the sum of the squares of the first one hundred natural
// numbers and the square of the sum.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	// run slow solution
	start := time.Now()
	res1, err := goeulerproject.SumOfSquares(1000000000)
	if err != nil {
		log.Fatal(err)
	}
	res2, err := goeulerproject.SquareOfSums(1000000000)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %v, elapsed (slow solution): %s\n",
		res2-res1, elapsed)

	// run fast solutions
	start = time.Now()
	res3, err := goeulerproject.FastSumOfSquares(1000000000)
	if err != nil {
		log.Fatal(err)
	}
	res4, err := goeulerproject.FastSquareOfSums(1000000000)
	if err != nil {
		log.Fatal(err)
	}
	elapsed = time.Since(start)
	fmt.Printf("result: %v, time elapsed (fast solution): %s\n",
		res4-res3, elapsed)
}
