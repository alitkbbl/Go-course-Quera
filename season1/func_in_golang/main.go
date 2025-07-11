package main

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	if x < 0 {
		return false
	}

	for i := 0; i*i <= x; i++ {
		if i*i == x {
			return true
		}
	}
	return false
}

func IsPalindrome(x int) bool {
	if x < 0 {
		x = -x
	}
	str := itoa(x)
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}
	digits := []byte{}
	n := num
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	result := []int{}
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map(input []int, m MapperFunc) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = m(v)
	}
	return result
}
