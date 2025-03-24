package finautils

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrorTooShortBase error = errors.New("too short base, min 3")
	ErrorTooLongBase  error = errors.New("too long base, max 19")
)

var factors = []int{7, 3, 1, 7, 3, 1, 7, 3, 1, 7, 3, 1, 7, 3, 1, 7, 3, 1, 7}

func ValidateFinnishReference(reference string) bool {
	base := reference[0 : len(reference)-1]
	genRef, err := CreateFinnishReference(base)
	if err != nil {
		return false
	}
	return reference == genRef
}

func CreateFinnishReference(base string) (string, error) {
	base = strings.TrimLeft(base, " 0")
	if len(base) > 19 {
		return "", ErrorTooLongBase
	}
	if len(base) < 3 {
		return "", ErrorTooShortBase
	}
	nums := make([]int, len(base))
	for i, ch := range base {
		nums[i] = int(ch) - 48
	}

	totalSum := 0
	j := 0
	for i := len(nums) - 1; i >= 0; i-- {
		totalSum += nums[i] * factors[j]
		j++
	}
	nextTen := getNextTenMultiplierVal(totalSum)
	check := nextTen - totalSum
	if check == 10 {
		check = 0
	}
	checkString := strconv.Itoa(check)
	return base + checkString, nil
}

func getNextTenMultiplierVal(base int) int {
	for base%10 != 0 {
		base++
	}
	return base
}
