SELECT
    toIPv6OrZero('2001:0db8:85a3:0000:0000:8a2e:0370:7334') AS valid_ipv6,
    toIPv6OrZero('invalid::ip') AS invalid_ipv6
