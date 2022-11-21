package utils

import (
	"fmt"
	"strconv"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MapArray[A any, B any](arr []A, fun func(a A) B) (res []B) {
	for _, elem := range arr {
		res = append(res, fun(elem))
	}
	return
}

func SprintOne(a any) string {
	return fmt.Sprint(a)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
