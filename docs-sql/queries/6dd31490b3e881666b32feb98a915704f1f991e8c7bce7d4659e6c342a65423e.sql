-- Over a non-HTTP interface (such as `clickhouse-client` or `clickhouse-local`) there are
-- no request headers, so the function returns an empty string. See the description above
-- for an HTTP example that returns the actual header value.
SELECT getClientHTTPHeader('Content-Type') SETTINGS allow_get_client_http_header = 1
