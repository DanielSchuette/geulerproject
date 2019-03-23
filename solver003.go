package goeulerproject

import "fmt"

// LargestPrimeFactor takes an integer and returns the largest
// prime factor of that number, see unit test in solver3_test.go
// for examples.
func LargestPrimeFactor(num int) (int, error) {
	// check if input is a prime number
	prime, err := IsPrimeNumber(num)
	if err != nil {
		return 0, fmt.Errorf("cannot determine primality of %d: %v", num, err)
	}
	if prime {
		return num, nil
	}

	// iterate over possible divisors of 'num', divide 'num' by
	// those divisors and test if the result is a prime number
	for i := 1; i < num; i++ {
		if (num % i) == 0 {
			// input is never malformed so errors are ignored
			p, _ := IsPrimeNumber(num / i)
			if p {
				return num / i, nil
			}
		}
	}
	return 0, fmt.Errorf("cannot find a prime factor")
}

// IsPrimeNumber takes an integer and returns a boolean indicating
// whether the input is a prime number or not or an error if the
// input is invalid.
func IsPrimeNumber(in int) (bool, error) {
	// check validity of input
	switch {
	case in < 1:
		return false, fmt.Errorf("invalid input: %d is less than 1", in)
	case in == 1:
		return false, nil
	}

	// iterate over all integers smaller than input and
	// return true if no modulo is 0
	for i := 2; i < in; i++ {
		if (in % i) == 0 {
			return false, nil
		}
	}
	return true, nil
}
