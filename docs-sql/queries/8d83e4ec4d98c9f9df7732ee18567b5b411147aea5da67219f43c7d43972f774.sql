INSERT INTO raw_events (id, payload) VALUES
(1, '{"type":"click","page":"/home"}'),
(2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

-- Confirm data round-trips intact
SELECT payload FROM raw_events WHERE id = 1

-- Verify you can still parse fields ad-hoc when needed
SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
