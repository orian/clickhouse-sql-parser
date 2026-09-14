SELECT
    toIPv4OrNull('192.168.1.1') AS valid_ip,
    toIPv4OrNull('invalid.ip') AS invalid_ip
