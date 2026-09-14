-- Insert sample data and inspect inferred paths
INSERT INTO events FORMAT JSONEachRow
{"timestamp":"2025-03-19 10:00:00","service":"api","level":"INFO","message":"request handled","host":"node-1","duration_ms":42,"attributes":{"http.status_code":200,"http.method":"GET","user.region":"eu-west","custom_tag":"abc"}}

SELECT JSONAllPathsWithTypes(attributes)
FROM events
FORMAT PrettyJSONEachRow
