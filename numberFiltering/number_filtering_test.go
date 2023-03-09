package numberFiltering

import (
	"reflect"
	"testing"
)

func TestEvenNumber(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	output := evenNumber(inputs)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}

func TestOddNumber(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 3, 5, 7, 9}
	output := oddNumber(inputs)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}

func TestPrimeNumber(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 3, 5, 7}
	output := primeNumber(inputs)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}

func TestOddPrimeNumber(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{3, 5, 7}
	output := oddPrimeNumber(inputs)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}

func TestEvenAndMultipleOfFive(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{10, 20}
	output := evenAndMultipleOfFive(inputs)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}

func TestOddAndMultipleOfThreeAndGreaterThanTen(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{15}
	output := oddAndMultipleOfThreeAndGreaterThanTen(inputs)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}

func TestFilterAll(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{9, 15}
	// odd, greater than 5, multiple of 3
	odd := func(n int) bool { return n%2 != 0 }
	greaterThan5 := func(n int) bool { return n > 5 }
	multipleOf3 := func(n int) bool { return n%3 == 0 }
	output := filterAll(inputs, odd, greaterThan5, multipleOf3)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}

func TestFilterAny(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}

	//prime := func(n int) bool { return isPrime(n) }
	greaterThan15 := func(n int) bool { return n > 15 }
	multipleOf5 := multipleOfN(5)
	output := filterAny(inputs, isPrime, greaterThan15, multipleOf5)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expected: %v but got: %v", expected, output)
	}
}
