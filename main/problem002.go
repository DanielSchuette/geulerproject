// Problem 2:
// Each new term in the Fibonacci sequence is generated
// by adding the previous two terms. By starting with 1
// and 2, the first 10 terms will be: 1, 2, 3, 5, 8, 13,
// 21, 34, 55, 89, ... By considering the terms in the
// Fibonacci sequence whose values do not exceed four
// million, find the sum of the even-valued terms.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	start := time.Now()
	res, err := goeulerproject.EvenFibs(4e6)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
