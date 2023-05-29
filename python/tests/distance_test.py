import pytest

import challenges.distance


def test_distance():
    a = challenges.distance.point(420, 69)
    b = challenges.distance.point(69, 420)
    assert challenges.distance.distance(a, b) == pytest.approx(496.39, 0.1)
