import challenges.fizz_buzz


def test_fizz_buzz():
    assert challenges.fizz_buzz.fizz_buzz(1) == "1"
    assert challenges.fizz_buzz.fizz_buzz(2) == "2"
    assert challenges.fizz_buzz.fizz_buzz(3) == "fizz"
    assert challenges.fizz_buzz.fizz_buzz(4) == "4"
    assert challenges.fizz_buzz.fizz_buzz(5) == "buzz"
    assert challenges.fizz_buzz.fizz_buzz(6) == "fizz"
    assert challenges.fizz_buzz.fizz_buzz(7) == "7"
    assert challenges.fizz_buzz.fizz_buzz(8) == "8"
    assert challenges.fizz_buzz.fizz_buzz(9) == "fizz"
    assert challenges.fizz_buzz.fizz_buzz(10) == "buzz"
    assert challenges.fizz_buzz.fizz_buzz(11) == "11"
    assert challenges.fizz_buzz.fizz_buzz(12) == "fizz"
    assert challenges.fizz_buzz.fizz_buzz(13) == "13"
    assert challenges.fizz_buzz.fizz_buzz(14) == "14"
    assert challenges.fizz_buzz.fizz_buzz(15) == "fizzbuzz"
    assert challenges.fizz_buzz.fizz_buzz(16) == "16"
    assert challenges.fizz_buzz.fizz_buzz(17) == "17"
    assert challenges.fizz_buzz.fizz_buzz(18) == "fizz"
    assert challenges.fizz_buzz.fizz_buzz(19) == "19"
    assert challenges.fizz_buzz.fizz_buzz(20) == "buzz"


def test_for_range():
    assert challenges.fizz_buzz.for_range(range(1, 21)) == {
        1: "1",
        2: "2",
        3: "fizz",
        4: "4",
        5: "buzz",
        6: "fizz",
        7: "7",
        8: "8",
        9: "fizz",
        10: "buzz",
        11: "11",
        12: "fizz",
        13: "13",
        14: "14",
        15: "fizzbuzz",
        16: "16",
        17: "17",
        18: "fizz",
        19: "19",
        20: "buzz",
    }
