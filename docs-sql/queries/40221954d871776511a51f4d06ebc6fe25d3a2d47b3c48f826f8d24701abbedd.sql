  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- 确认数据能够完整往返
  SELECT payload FROM raw_events WHERE id = 1

  -- 验证在需要时仍可临时解析字段
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
