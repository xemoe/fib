package fib

//
// @author Teerapong Ladlee
// @created_at 2019-03-25
//

import (
	"testing"
	"reflect"
)

func TestFib(t *testing.T) {

	var result string

	result = Fib(1)
	if result != "1" {
		t.Errorf("Fib(1) should return 1, %q is returned", result)
	}

	result = Fib(2)
	if result != "2" {
		t.Errorf("Fib(2) should return 1, %q is returned", result)
	}

	result = Fib(3)
	if result != "fizz" {
		t.Errorf("Fib(3) should return fizz, %q is returned", result)
	}

	result = Fib(5)
	if result != "buzz" {
		t.Errorf("Fib(5) should return buzz, %q is returned", result)
	}

	result = Fib(10)
	if result != "buzz" {
		t.Errorf("Fib(10) should return buzz, %q is returned", result)
	}

	result = Fib(15)
	if result != "fizzbuzz" {
		t.Errorf("Fib(15) should return fizzbuzz, %q is returned", result)
	}
}

func TestFibSequence(t *testing.T) {

	var result []string
	var expected []string

	expected = []string{"1"}
	result, _ = FibSequence([]int{1})

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FibSequence([]int{1}) should return expected slice, %q is returned", result)
	}

	expected = []string{"1", "2", "fizz", "4", "buzz", "buzz", "fizzbuzz",}
	result, _ = FibSequence([]int{1, 2, 3, 4, 5, 10, 15})

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FibSequence([]int{1}) should return expected slice, %q is returned", result)
	}
}

func TestFibSequence_WithEmptyInputSlice(t *testing.T) {

	var result []string
	var err error

	result, err = FibSequence([]int{})

	if result != nil {
		t.Errorf("FibSequence([]int{}) should return expected nil, %q is returned", result)
	}
	if err == nil {
		t.Errorf("FibSequence([]int{}) returned error should not be nil, %q is returned", err)
	}
}
