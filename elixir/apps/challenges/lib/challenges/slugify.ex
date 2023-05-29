defmodule Challenges.Slugify do
  @spec slugify(String.t()) :: String.t()
  def slugify(str) do
    str
    |> String.downcase()
    |> String.replace(~r/['"!@#$%^&\*\(\)\[\]\{\};:\,\.\/<>\?\|`~=_+ ]/, "-")
    |> String.trim("-")
  end

  @spec deslugify(String.t()) :: String.t()
  def deslugify(str) do
    String.split(str, ~r/\-/)
    |> Enum.map(fn s -> String.capitalize(s) end)
    |> Enum.join(" ")
  end
end
