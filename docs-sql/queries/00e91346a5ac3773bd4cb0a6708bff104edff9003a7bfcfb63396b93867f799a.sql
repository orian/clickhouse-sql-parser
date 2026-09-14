WITH
    '192.168.1.1' AS valid_IPv4_string,
    '999.999.999.999' AS invalid_IPv4_string,
    'not_an_ip' AS malformed_string
SELECT
    toIPv4OrDefault(valid_IPv4_string) AS valid,
    toIPv4OrDefault(invalid_IPv4_string) AS default_value,
    toIPv4OrDefault(malformed_string, toIPv4('8.8.8.8')) AS provided_default;
