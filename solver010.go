package goeulerproject

import (
	"fmt"
	"math"
)

// SumOfPrimes takes an integer n as its input and returns
// the sum of all prime numbers smaller than n. See problem10.go
// for a more detailed description of the problem.
func SumOfPrimes(n int) (int, error) {
	// test validity of input
	if n < 2 {
		return 0, fmt.Errorf("invalid input: %d is less than 2", n)
	}

	// find and sum the first n primes and return the result
	var sum, counter int
	currentNum := 2
	for {
		if IsPrime(currentNum) {
			counter++
			sum += currentNum
		}
		currentNum++
		if currentNum > n {
			return sum, nil
		}
	}
}

// IsPrime returns true if the input is a prime number.
func IsPrime(n int) bool {
	// test corner cases
	switch n {
	case 1:
		return false
	case 2:
		return true
	}

	// only odd numbers are prime
	if n%2 == 0 {
		return false
	}

	for i := 3; i < (int(math.Sqrt(float64(n))) + 1); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}
