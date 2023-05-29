defmodule Challenges.Distance do
  @type point() :: {integer(), integer()}

  @spec distance(point(), point()) :: float()
  def distance(a, b) do
    {{ax, ay}, {bx, by}} = {a, b}
    :math.sqrt(((ax - bx)**2) + ((ay - by)**2))
  end
end
