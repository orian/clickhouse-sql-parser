SELECT count()
FROM otel_logs_index_body
WHERE hasAllTokens(Body, ['Connection', 'accepted'])
