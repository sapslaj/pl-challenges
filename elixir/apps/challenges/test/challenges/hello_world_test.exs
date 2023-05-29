defmodule Challenges.HelloWorldTest do
  use ExUnit.Case, async: true

  test "says hello world" do
    assert Challenges.HelloWorld.hello_world() == "hello world"
  end

  test "prints hello world" do
    assert ExUnit.CaptureIO.capture_io(fn ->
      Challenges.HelloWorld.print()
    end) == "hello world\n"
  end
end
