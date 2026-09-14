ALTER TABLE otel_logs
ADD INDEX idx_log_attr_items LogAttributeItems
TYPE text(tokenizer = 'array')
