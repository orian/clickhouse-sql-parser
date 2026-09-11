SELECT
    IPv6NumToString(IPv6StringToNumOrDefault('2001:db8::1')) AS valid,
    IPv6NumToString(IPv6StringToNumOrDefault('invalid')) AS invalid;
