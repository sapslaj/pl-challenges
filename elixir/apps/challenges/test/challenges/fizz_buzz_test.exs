defmodule Challenges.FizzBuzzTest do
  use ExUnit.Case, async: true

  test "correct outputs" do
    assert Challenges.FizzBuzz.fizz_buzz(1) == "1"
    assert Challenges.FizzBuzz.fizz_buzz(2) == "2"
    assert Challenges.FizzBuzz.fizz_buzz(3) == "fizz"
    assert Challenges.FizzBuzz.fizz_buzz(4) == "4"
    assert Challenges.FizzBuzz.fizz_buzz(5) == "buzz"
    assert Challenges.FizzBuzz.fizz_buzz(6) == "fizz"
    assert Challenges.FizzBuzz.fizz_buzz(7) == "7"
    assert Challenges.FizzBuzz.fizz_buzz(8) == "8"
    assert Challenges.FizzBuzz.fizz_buzz(9) == "fizz"
    assert Challenges.FizzBuzz.fizz_buzz(10) == "buzz"
    assert Challenges.FizzBuzz.fizz_buzz(11) == "11"
    assert Challenges.FizzBuzz.fizz_buzz(12) == "fizz"
    assert Challenges.FizzBuzz.fizz_buzz(13) == "13"
    assert Challenges.FizzBuzz.fizz_buzz(14) == "14"
    assert Challenges.FizzBuzz.fizz_buzz(15) == "fizzbuzz"
    assert Challenges.FizzBuzz.fizz_buzz(16) == "16"
    assert Challenges.FizzBuzz.fizz_buzz(17) == "17"
    assert Challenges.FizzBuzz.fizz_buzz(18) == "fizz"
    assert Challenges.FizzBuzz.fizz_buzz(19) == "19"
    assert Challenges.FizzBuzz.fizz_buzz(20) == "buzz"
  end

  test "range" do
    assert Challenges.FizzBuzz.range(1..20) == %{
      1 => "1",
      2 => "2",
      3 => "fizz",
      4 => "4",
      5 => "buzz",
      6 => "fizz",
      7 => "7",
      8 => "8",
      9 => "fizz",
      10 => "buzz",
      11 => "11",
      12 => "fizz",
      13 => "13",
      14 => "14",
      15 => "fizzbuzz",
      16 => "16",
      17 => "17",
      18 => "fizz",
      19 => "19",
      20 => "buzz",
    }
  end
end
