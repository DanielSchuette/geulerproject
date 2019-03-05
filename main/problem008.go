// Problem 8:
// The four adjacent digits in the 1000-digit number
// that have the greatest product are 9 × 9 × 8 × 9 =
// 5832. Find the thirteen adjacent digits in the 1000-
// digit number that have the greatest product. What is
// the value of this product?
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.FindGreatestProduct(13)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
