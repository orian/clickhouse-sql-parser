  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- Убедитесь, что данные записываются и читаются без изменений
  SELECT payload FROM raw_events WHERE id = 1

  -- Убедитесь, что при необходимости поля всё ещё можно разбирать на лету
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
