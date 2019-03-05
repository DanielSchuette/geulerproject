// Problem 4:
// A palindromic number reads the same both ways. The largest
// palindrome made from the product of two 2-digit numbers is
// 9009 = 91 × 99. Find the largest palindrome made from the
// product of two 3-digit numbers.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.FindPalindrome(3)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
