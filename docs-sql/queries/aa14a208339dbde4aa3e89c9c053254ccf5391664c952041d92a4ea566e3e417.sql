ALTER TABLE default.otel_logs
MODIFY TTL TimestampTime + toIntervalDay(7);
