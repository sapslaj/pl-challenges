defmodule Challenges.SlugifyTest do
  use ExUnit.Case, async: true

  test "slugify" do
    assert Challenges.Slugify.slugify("My Test Title") == "my-test-title"
    assert Challenges.Slugify.slugify("String w/ some extra &stuff'") == "string-w--some-extra--stuff"
  end

  test "deslugify" do
    assert Challenges.Slugify.deslugify("my-test-title") == "My Test Title"
  end
end
