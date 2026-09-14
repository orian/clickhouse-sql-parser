ALTER TABLE otel_logs ADD INDEX idx_body Body TYPE text(tokenizer = splitByNonAlpha) GRANULARITY 100000000;
ALTER TABLE otel_logs MATERIALIZE INDEX idx_body;
