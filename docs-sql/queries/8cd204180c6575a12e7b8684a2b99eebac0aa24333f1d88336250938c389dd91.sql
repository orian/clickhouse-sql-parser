CREATE TABLE tickets (body String) ENGINE = Memory;
INSERT INTO tickets VALUES ('Contact Jane Doe at jane@example.com.');
SELECT aiRedact(body, [], map('replacement', '***')) FROM tickets LIMIT 5
