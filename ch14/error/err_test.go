package error_test

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var GreaterThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	// if n < 2 || n > 100 {
	// 	return nil, errors.New("n should be in [2,100]")
	// }

	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, GreaterThanHundredError
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if list, err := GetFibonacci(-5); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number.")
		}

		if err == GreaterThanHundredError {
			t.Error("Need a less number.")
		}

	} else {
		t.Log(list)
	}

}
