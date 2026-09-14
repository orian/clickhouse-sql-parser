  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- データが損なわれずに往復変換できることを確認
  SELECT payload FROM raw_events WHERE id = 1

  -- 必要なときにその場でフィールドをパースできることを確認
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
