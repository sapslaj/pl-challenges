package fizzbuzz_test

import (
	"testing"

	fizzbuzz "github.com/sapslaj/pl-challenges/go/challenges/fizz_buzz"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "1", fizzbuzz.FizzBuzz(1))
	assert.Equal(t, "2", fizzbuzz.FizzBuzz(2))
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(3))
	assert.Equal(t, "4", fizzbuzz.FizzBuzz(4))
	assert.Equal(t, "buzz", fizzbuzz.FizzBuzz(5))
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(6))
	assert.Equal(t, "7", fizzbuzz.FizzBuzz(7))
	assert.Equal(t, "8", fizzbuzz.FizzBuzz(8))
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(9))
	assert.Equal(t, "buzz", fizzbuzz.FizzBuzz(10))
	assert.Equal(t, "11", fizzbuzz.FizzBuzz(11))
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(12))
	assert.Equal(t, "13", fizzbuzz.FizzBuzz(13))
	assert.Equal(t, "14", fizzbuzz.FizzBuzz(14))
	assert.Equal(t, "fizzbuzz", fizzbuzz.FizzBuzz(15))
	assert.Equal(t, "16", fizzbuzz.FizzBuzz(16))
	assert.Equal(t, "17", fizzbuzz.FizzBuzz(17))
	assert.Equal(t, "fizz", fizzbuzz.FizzBuzz(18))
	assert.Equal(t, "19", fizzbuzz.FizzBuzz(19))
	assert.Equal(t, "buzz", fizzbuzz.FizzBuzz(20))
}

func TestRange(t *testing.T) {
	t.Parallel()

	expected := map[int]string{
		1:  "1",
		2:  "2",
		3:  "fizz",
		4:  "4",
		5:  "buzz",
		6:  "fizz",
		7:  "7",
		8:  "8",
		9:  "fizz",
		10: "buzz",
		11: "11",
		12: "fizz",
		13: "13",
		14: "14",
		15: "fizzbuzz",
		16: "16",
		17: "17",
		18: "fizz",
		19: "19",
		20: "buzz",
	}
	assert.Equal(t, expected, fizzbuzz.Range(1, 20))
}
