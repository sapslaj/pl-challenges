import challenges.slugify


def test_slugify():
    assert challenges.slugify.slugify("My Test Title") == "my-test-title"
    assert challenges.slugify.slugify("String w/ some extra &stuff'") == "string-w--some-extra--stuff"


def test_deslugify():
    assert challenges.slugify.deslugify("my-test-title") == "My Test Title"
