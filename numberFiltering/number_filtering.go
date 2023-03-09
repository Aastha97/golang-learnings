package numberFiltering

import (
	"math"
)

func evenNumber(num []int) []int {
	var output []int
	for _, n := range num {
		if isEven(n) {
			output = append(output, n)
		}
	}
	return output
}

func oddNumber(num []int) []int {
	var output []int
	for _, n := range num {
		if !isEven(n) {
			output = append(output, n)
		}
	}
	return output
}

func primeNumber(num []int) []int {
	var output []int
	for _, n := range num {
		if isPrime(n) {
			output = append(output, n)
		}
	}
	return output
}
func isPrime(n int) bool {
	for n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func oddPrimeNumber(num []int) []int {
	var output []int
	for _, n := range num {
		if isPrime(n) && !isEven(n) {
			output = append(output, n)
		}
	}
	return output
}

func evenAndMultipleOfFive(num []int) []int {
	var output []int
	for _, n := range num {
		if isEven(n) && multipleOf(n, 5) {
			output = append(output, n)
		}
	}
	return output
}

func oddAndMultipleOfThreeAndGreaterThanTen(num []int) []int {
	var output []int
	for _, n := range num {
		if !isEven(n) && multipleOf(n, 3) && n > 10 {
			output = append(output, n)
		}
	}
	return output
}

func isEven(n int) bool {
	return n%2 == 0
}

func multipleOf(n, multiple int) bool {
	return n%multiple == 0
}

func multipleOfN(n int) func(num int) bool {
	return func(num int) bool {
		return num%n == 0
	}
}

type Condition func(n int) bool

func filterAll(nums []int, conditions ...Condition) []int {
	var output []int
	for _, n := range nums {
		var flag bool = true
		for _, con := range conditions {
			if !con(n) {
				flag = false
				break
			}
		}
		if flag {
			output = append(output, n)
		}
	}
	return output
}

func filterAny(nums []int, conditions ...Condition) []int {
	var output []int
	for _, n := range nums {
		for _, con := range conditions {
			if con(n) {
				output = append(output, n)
				break
			}
		}
	}
	return output
}
