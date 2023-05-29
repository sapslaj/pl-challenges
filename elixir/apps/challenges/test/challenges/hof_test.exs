defmodule Challenges.HofTest do
  use ExUnit.Case, async: true

  defp list do
    ["abc", "def"]
  end

  test "map" do
    result = ["ABC", "DEF"]
    assert(Enum.map(list(), fn x -> String.upcase(x) end) == result)
  end

  test "filter" do
    result = ["abc"]
    assert(Enum.filter(list(), fn x -> String.contains?(x, "a") end) == result)
  end

  test "reduce" do
    result = "abcdef"
    assert(Enum.reduce(list(), "", fn x, acc -> acc <> x end) == result)
  end

  test "reverse" do
    result = ["def", "abc"]
    assert(Enum.reverse(list()) == result)
  end

  test "all" do
    assert list()
    |> Enum.map(fn x -> String.length(x) > 0 end)
    |> Enum.all?()
  end

  test "any" do
    assert list()
    |> Enum.map(fn x -> String.contains?(x, "a") end)
    |> Enum.any?()
  end
end
