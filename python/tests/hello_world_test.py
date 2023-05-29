import challenges.hello_world


def test_hello_world():
    assert challenges.hello_world.hello_world() == "hello world"


def test_puts(capsys):
    challenges.hello_world.puts()
    captured = capsys.readouterr()
    assert captured.out == "hello world\n"
