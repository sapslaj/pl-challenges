package fizzbuzz

import "core:fmt"

fizz_buzz :: proc(n: int) -> string {
  if n % 3 == 0 && n % 5 == 0 {
    return "fizzbuzz"
  } else if n % 3 == 0 {
    return "fizz"
  } else if n % 5 == 0 {
    return "buzz"
  } else {
    return fmt.aprint(n)
  }
}

fizzbuzz_range :: proc(from: int, to: int) -> map[int]string {
  result := make(map[int]string)
  for n := from; n <= to ; n += 1 {
    result[n] = fizz_buzz(n)
  }
  return result
}

main :: proc() {
  for n, res in fizzbuzz_range(1, 20) {
    fmt.printf("[%d]: %s\n", n, res)
  }
}
