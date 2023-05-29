from typing import Iterable
import ipaddress


def categorize_ipv4_ipv6(l: Iterable[str]) -> tuple[Iterable[str], Iterable[str]]:
    ipv4s = []
    ipv6s = []
    for address in l:
        match ipaddress.ip_address(address).version:
            case 4:
                ipv4s.append(address)
            case 6:
                ipv6s.append(address)
    return ipv4s, ipv6s
