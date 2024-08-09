package utils

import "strconv"

func Filter[T any](arr []T, f func(T) bool) []T {
	var result []T
	for _, value := range arr {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func MapStringArrayToInt(arr []string) []int {
	return Map(arr, func(value string) int {
		i, _ := strconv.Atoi(value)
		return i
	})
}

func MapStringArrayToInt64(arr []string) []int64 {
	return Map(arr, func(value string) int64 {
		i, _ := strconv.ParseInt(value, 10, 64)
		return i
	})
}

func Map[T any, R any](arr []T, f func(T) R) []R {
	var result []R
	for _, value := range arr {
		result = append(result, f(value))
	}
	return result
}

// GetValueOrKey Returns the value of the key in the map if it exists, otherwise returns the key itself.
func GetValueOrKey[T comparable](mapp map[T]T, key T) T {
	return GetOr(mapp, key, key)
}

func GetOr[T comparable, R any](mapp map[T]R, key T, defaultValue R) R {
	if value, present := mapp[key]; present {
		return value
	} else {
		return defaultValue
	}
}

func CreateIntArray(start int, end int) []int {
	var arr []int
	for i := start; i <= end; i++ {
		arr = append(arr, i)
	}
	return arr
}

func ArraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
