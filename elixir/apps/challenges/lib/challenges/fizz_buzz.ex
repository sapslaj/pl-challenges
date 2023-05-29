defmodule Challenges.FizzBuzz do
  @spec fizz_buzz(integer()) :: String.t()
  def fizz_buzz(n) do
    cond do
      rem(n, 3) == 0 and rem(n, 5) == 0 -> "fizzbuzz"
      rem(n, 3) == 0 -> "fizz"
      rem(n, 5) == 0 -> "buzz"
      true -> to_string(n)
    end
  end

  @spec range(Range.t()) :: map()
  def range(r) do
    Enum.reduce(r, %{}, fn n, m ->
      update_in(m[n], fn _ -> fizz_buzz(n) end)
    end)
  end
end
