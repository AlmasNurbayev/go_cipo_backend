package utils

import (
	"strconv"
	"strings"
)

func String2Int64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func String2Int(str string) int {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int(i)
}

func String2ArrayInt64(str string, divider string) []int64 {
	arr := make([]int64, 0)
	parts := strings.Split(str, divider)
	for _, v := range parts {
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			continue
		}
		arr = append(arr, i)
	}
	return arr
}

func String2Float32(str string) float32 {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(f)
}
