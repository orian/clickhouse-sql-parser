  INSERT INTO raw_events (id, payload) VALUES
  (1, '{"type":"click","page":"/home"}'),
  (2, '{"type":"purchase","item":"SKU-99","amount":49.99}')

  -- Confirme que os dados permanecem intactos na ida e volta
  SELECT payload FROM raw_events WHERE id = 1

  -- Verifique que ainda é possível extrair campos de forma ad hoc quando necessário
  SELECT JSONExtractString(payload, 'type') AS event_type FROM raw_events
