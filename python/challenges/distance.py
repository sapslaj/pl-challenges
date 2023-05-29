import math

Point = tuple[int, int]


def point(x: int, y: int) -> Point:
    return Point([x, y])


def distance(a: Point, b: Point) -> float:
    return math.sqrt(((a[0] - b[0]) ** 2) + ((a[1] - b[1]) ** 2))
