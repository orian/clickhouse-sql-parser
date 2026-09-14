  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- 데이터가 손상 없이 그대로 저장·조회되는지 확인
  SELECT payload FROM raw_events WHERE id = 1

  -- 필요할 때는 여전히 즉석에서 필드를 파싱할 수 있는지 확인
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
