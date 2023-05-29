defmodule Challenges.DistanceTest do
  use ExUnit.Case, async: true

  test "distance" do
    a = {420, 69}
    b = {69, 420}
    assert_in_delta(Challenges.Distance.distance(a, b), 496.39, 0.1)
  end
end
