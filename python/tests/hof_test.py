import functools

l = ["abc", "def"]


def test_map():
    result = ["ABC", "DEF"]
    assert [x.upper() for x in l] == result
    assert list(map(lambda x: x.upper(), l)) == result


def test_filter():
    result = ["abc"]
    assert [x for x in l if "a" in x] == result
    assert list(filter(lambda x: "a" in x, l)) == result


def test_reduce():
    result = "abcdef"
    assert functools.reduce(lambda r, x: "".join([r, x]), l, "") == result
    "".join(l) == result


def test_reverse():
    result = ["def", "abc"]
    assert list(reversed(l)) == result


def test_all():
    assert all([len(x) > 0 for x in l])


def test_any():
    assert any(["a" in x for x in l])
