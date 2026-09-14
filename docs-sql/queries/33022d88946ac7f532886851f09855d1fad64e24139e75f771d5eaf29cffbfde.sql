SELECT
    IPv4StringToNumOrDefault('127.0.0.1') AS valid,
    IPv4StringToNumOrDefault('invalid') AS invalid;
