// Problem 9:
// A Pythagorean triplet is a set of three natural numbers,
// a < b < c, for which, a^2 + b^2 = c^2. For example,
// 3^2 + 4^2 = 9 + 16 = 25 = 5^2. There exists exactly one
// Pythagorean triplet for which a + b + c = 1000. Find the
// product a*b*c.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	a, b, c, err := goeulerproject.PythagoreanTriplet(1000)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("a: %v, b: %v, c: %v, sum: %v, product: %v, ",
		a, b, c, a+b+c, a*b*c)
	fmt.Printf("elapsed: %s\n", elapsed)
}
