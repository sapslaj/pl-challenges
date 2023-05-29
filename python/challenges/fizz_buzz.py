def fizz_buzz(n: int) -> str:
    if n % 3 == 0 and n % 5 == 0:
        return "fizzbuzz"
    elif n % 3 == 0:
        return "fizz"
    elif n % 5 == 0:
        return "buzz"
    else:
        return str(n)


def for_range(r: range) -> dict[int, str]:
    result = {}
    for n in r:
        result[n] = fizz_buzz(n)
    return result
