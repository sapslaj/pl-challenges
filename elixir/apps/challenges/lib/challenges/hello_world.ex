defmodule Challenges.HelloWorld do
  def hello_world() do
    "hello world"
  end

  def print() do
    IO.puts(hello_world())
  end
end
