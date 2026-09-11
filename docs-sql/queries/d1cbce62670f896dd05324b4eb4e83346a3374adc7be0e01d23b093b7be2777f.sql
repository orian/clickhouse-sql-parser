-- Ignore an index by name
SELECT * FROM logs
WHERE hasToken(lower(msg), 'exception')
SETTINGS ignore_data_skipping_indices = 'msg_token';
