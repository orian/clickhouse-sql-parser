SELECT
    toIPv4OrZero('192.168.1.1') AS valid_ip,
    toIPv4OrZero('invalid.ip') AS invalid_ip
