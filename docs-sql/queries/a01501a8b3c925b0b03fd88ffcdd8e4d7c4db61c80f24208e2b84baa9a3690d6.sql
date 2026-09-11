SELECT count()
FROM otel_logs
WHERE hasAllTokens(Body, ['Connection', 'accepted'])
