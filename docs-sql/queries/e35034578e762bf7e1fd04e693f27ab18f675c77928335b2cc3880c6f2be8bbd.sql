SELECT count()
FROM otel_logs_text_body_preprocessed
WHERE hasAllTokens(Body, ['Connection', 'accepted'])
