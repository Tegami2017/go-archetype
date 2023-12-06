package util

import (
	"strconv"
	"strings"
)

type ArrayUtil struct{}

var Array = &ArrayUtil{}

func (r ArrayUtil) ToIntArr(param string, separator string) []int {
	if len(param) == 0 {
		return make([]int, 0)
	}
	strList := strings.Split(param, separator)
	if len(strList) == 0 {
		return make([]int, 0)
	}
	intList := make([]int, len(strList))
	for i, v := range strList {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			panic("invalid item")
		}
		intList[i] = intVal
	}
	return intList
}
