package goeulerproject

import (
	"errors"
	"testing"
)

type testCase12 struct {
	in   int
	want int
	err  error
}

func TestSumOfPrimes(t *testing.T) {
	cases := []testCase12{
		// test invalid input
		{
			in:   0,
			want: 0,
			err:  errors.New("invalid input: 0 is less than 2"),
		},
		// test valid example
		{
			in:   12,
			want: 28,
			err:  nil,
		},
	}

	// iterate over test cases
	for _, c := range cases {
		got, err := SumOfPrimes(c.in)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("SumOfPrimes(%d) == %v, want %v\n",
				c.in, got, c.want)
		}

		// if no error is returned, test if none is expected
		if err == nil && c.err != nil {
			t.Errorf("SumOfPrimes(%v) returned error %v, want %v\n",
				c.in, err, c.err)
		}

		// if error is returned, test if an error is expected
		if err != nil {
			// if c.err is nil, print wanted and received error
			// else if an error is wanted and received but error
			// messages are not the same
			// print wanted and received error
			if c.err == nil {
				t.Errorf("SumOfPrimes(%v) returned error %v, want %v\n",
					c.in, err, c.err)
			} else if err.Error() != c.err.Error() {
				t.Errorf("SumOfPrimes(%v) returned error %v, want %v\n",
					c.in, err, c.err)
			}
		}
	}
}

type testCase13 struct {
	in   int
	want bool
	err  error
}

func TestIsPrime(t *testing.T) {
	cases := []testCase13{
		{
			in:   1,
			want: false,
		},
		{
			in:   2,
			want: true,
		},
		{
			in:   4,
			want: false,
		},
		{
			in:   7,
			want: true,
		},
		{
			in:   9,
			want: false,
		},
		{
			in:   15,
			want: false,
		},
	}

	// iterate over test cases
	for _, c := range cases {
		got := IsPrime(c.in)

		// test similarity of expected and received value
		if got != c.want {
			t.Errorf("IsPrime(%d) == %v, want %v\n",
				c.in, got, c.want)
		}
	}
}
