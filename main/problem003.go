// Problem 3:
// The prime factors of 13195 are 5, 7, 13 and 29. What is
// the largest prime factor of the number 600851475143?
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/DanielSchuette/goeulerproject"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("provide the number for which to find the largest prime factor as an argument")
	}
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	res, err := goeulerproject.LargestPrimeFactor(input) // 600851475143
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("result: %d, elapsed: %s\n", res, elapsed)
}
