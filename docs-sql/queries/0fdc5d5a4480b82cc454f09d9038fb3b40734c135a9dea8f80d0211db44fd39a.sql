SELECT
    IPv6NumToString(IPv6StringToNumOrNull('2001:db8::1')) AS valid,
    IPv6StringToNumOrNull('invalid') AS invalid;
