package fizzbuzz

import "strconv"

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "fizzbuzz"
	} else if n%3 == 0 {
		return "fizz"
	} else if n%5 == 0 {
		return "buzz"
	} else {
		return strconv.Itoa(n)
	}
}

func Range(from, to int) map[int]string {
	result := make(map[int]string)
	for i := from; i <= to; i++ {
		result[i] = FizzBuzz(i)
	}
	return result
}
