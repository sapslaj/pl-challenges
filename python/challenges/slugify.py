import re


def slugify(s: str) -> str:
    return re.sub(r"['\"!@#$%^&\*\(\)\[\]\{\};:\,\./<>\?\|`~=_+ ]", "-", s.lower()).strip("-")


def deslugify(s: str) -> str:
    return " ".join([part.capitalize() for part in s.split("-")])
