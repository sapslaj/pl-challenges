import challenges.ip_addr


def test_categorize_ipv4_ipv6():
    ips = [
        "::1",
        "0.0.0.0",
        "192.168.0.1",
        "fe80::1",
    ]
    ipv4 = [
        "0.0.0.0",
        "192.168.0.1",
    ]
    ipv6 = [
        "::1",
        "fe80::1",
    ]
    assert challenges.ip_addr.categorize_ipv4_ipv6(ips) == (ipv4, ipv6)
