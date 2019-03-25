package fib

//
// @author Teerapong Ladlee
// @created_at 2019-03-25
//

import (
	"strconv"
	"fmt"
)

const (
	REASON_FIZZ = 3
	REASON_BUZZ = 5
	WORD_FIZZ = "fizz"
	WORD_BUZZ = "buzz"
)

func Fib(num int) string {
	var result string

	if num % REASON_FIZZ == 0 {
		result += WORD_FIZZ
	}

	if num % REASON_BUZZ == 0 {
		result += WORD_BUZZ
	}

	if result == "" {
		result = strconv.Itoa(num)
	}

	return result
}

func FibSequence(numsl []int) ([]string, error) {
	var result []string
	var err error

	if len(numsl) == 0 {
		err = fmt.Errorf("Input slices should not be empty")
	}

	for _, num := range numsl {
		result = append(result, Fib(num))
	}

	return result, err
}
