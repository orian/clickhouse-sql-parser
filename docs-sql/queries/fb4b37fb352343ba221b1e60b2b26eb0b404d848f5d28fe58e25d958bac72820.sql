WITH
    '2001:0db8:85a3:0000:0000:8a2e:0370:7334' AS valid_IPv6_string,
    '2001:0db8:85a3::8a2e:370g:7334' AS invalid_IPv6_string,
    'not_an_ipv6' AS malformed_string
SELECT
    toIPv6OrDefault(valid_IPv6_string) AS valid,
    toIPv6OrDefault(invalid_IPv6_string) AS default_value,
    toIPv6OrDefault(malformed_string, toIPv6('::1')) AS provided_default;
