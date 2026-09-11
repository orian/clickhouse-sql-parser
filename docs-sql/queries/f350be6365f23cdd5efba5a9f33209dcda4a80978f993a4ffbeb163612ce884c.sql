-- Confirm column types match expectations
DESCRIBE TABLE events FORMAT Vertical

-- Insert and query to validate the schema handles your data
INSERT INTO events FORMAT JSONEachRow
{"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42}

SELECT service, level, duration_ms FROM events WHERE service = 'api'
